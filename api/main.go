package main

// Import the required packages and functions
import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"golang.org/x/crypto/bcrypt"

	wkhtml "github.com/SebastiaanKlippert/go-wkhtmltopdf"
)

var db *sql.DB

type User struct {
	ID             int64
	Jobtitle       string
	Firstname      string
	Lastname       string
	Email          string
	Phone          string
	Address        string
	City           string
	Country        string
	Postalcode     string
	Dateofbirth    string
	Nationality    string
	Summary        string
	Workexperience string
	Education      string
	Skills         string
	Languages      string
}

type Template struct {
	Id   int64
	Path string
}

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64), // this key is used for signing
	securecookie.GenerateRandomKey(32), // this key is used for encryption
)

func verifyLogin(username, password string) bool {
	var hashedPassword string
	err := db.QueryRow("SELECT password FROM userlogin WHERE username = ?", username).Scan(&hashedPassword)
	if err != nil {
		return false
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

func setSession(userName string, w http.ResponseWriter) {
	value := map[string]string{
		"name": userName,
	}

	if encoded, err := cookieHandler.Encode("session", value); err == nil {
		cookie := &http.Cookie{
			Name:  "session",
			Value: encoded,
			Path:  "/",
		}
		http.SetCookie(w, cookie)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	username := strings.TrimSpace(r.Form.Get("username"))
	password := r.Form.Get("password")

	response := struct {
		Username string `json:"username"`
	}{
		Username: username,
	}

	if verifyLogin(username, password) {
		w.Header().Set("Content-Type", "application/json")
		setSession(username, w)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	}

}

func logoutHandler(w http.ResponseWriter, r *http.Request) {
	//clear session
	cookie := &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User logout successfully"))
}

func signupHandler(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	email := strings.TrimSpace(r.Form.Get("email"))
	password := r.Form.Get("password")

	if email == "" || password == "" {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = db.Exec("INSERT INTO userlogin (username, password) VALUES (?,?)", email, hashedPassword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	w.Write([]byte("User signup successfully"))
}

func homeUsers(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	var users []User
	rows, err := db.Query("SELECT id, firstname, lastname, email FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Firstname, &user.Lastname, &user.Email)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	rows, err := db.Query("SELECT id, jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

func showUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	var user User
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", 1)
	if err := row.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages); err != nil {
		if err == sql.ErrNoRows {
			http.NotFound(w, r)
			return
		}
		http.Error(w, fmt.Sprintf("Error fetching user data: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func createUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	var user User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Insert the user into the database
	_, err := db.Exec("INSERT INTO users (Jobtitle, Firstname, Lastname, Email, Phone, Address, City, Country, Postalcode, Dateofbirth, Nationality, Summary, Workexperience, Education, Skills, Languages) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)",
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte("User created successfully"))
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "PUT" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	var user User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	// Check if the user exists
	_, err = db.Exec("UPDATE users SET jobtitle=?, first_name=?, last_name=?, email=?, phone=?, address=?, city=?, country=?, postal_code=?, date_of_birth=?, nationality=?, summary=?, work_experience=?, education=?, skills=?, languages=? WHERE id=?",
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.ID = int64(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "DELETE" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	userID := r.URL.Query().Get("id")
	if userID == "" {
		http.NotFound(w, r)
		return
	}

	id, err := strconv.ParseInt(userID, 10, 64)
	if err != nil || id < 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "ID should be an integer")
		return
	}

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error preparing delete statement: %v", err)
		return
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error executing delete statement: %v", err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

func generateTemplate(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	template_id := query["template"]
	user_id := query["user"]

	iduser_int, err := strconv.Atoi(user_id[0])
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}

	idtemplate_int, err := strconv.Atoi(template_id[0])
	if err != nil {
		log.Printf("An error occurred: %v", err)
	}

	var user User
	var template Template

	// Get the path of the template
	row1 := db.QueryRow("SELECT Path FROM template WHERE id = ?", idtemplate_int)
	row1.Scan(&template.Path)
	row := db.QueryRow("SELECT * FROM users WHERE id = ?", iduser_int)
	row.Scan(&user.ID, &user.Jobtitle, &user.Firstname, &user.Lastname, &user.Email, &user.Phone, &user.Address, &user.City, &user.Country, &user.Postalcode, &user.Dateofbirth, &user.Nationality, &user.Summary, &user.Workexperience, &user.Education, &user.Skills, &user.Languages)

	
	htmlContent, err := os.ReadFile(template.Path)
	if err != nil {
		panic(err)
	}

	htmlString := string(htmlContent)

	htmlString = strings.ReplaceAll(htmlString, "{{Firstname}}", user.Firstname)
	htmlString = strings.ReplaceAll(htmlString, "{{Lastname}}", user.Lastname)
	htmlString = strings.ReplaceAll(htmlString, "{{Jobtitle}}", user.Jobtitle)
	htmlString = strings.ReplaceAll(htmlString, "{{Email}}", user.Email)
	htmlString = strings.ReplaceAll(htmlString, "{{Phone}}", user.Phone)
	htmlString = strings.ReplaceAll(htmlString, "{{Address}}", user.Address)
	htmlString = strings.ReplaceAll(htmlString, "{{City}}", user.City)
	htmlString = strings.ReplaceAll(htmlString, "{{Country}}", user.Country)
	htmlString = strings.ReplaceAll(htmlString, "{{Postalcode}}", user.Postalcode)
	htmlString = strings.ReplaceAll(htmlString, "{{Dateofbirth}}", user.Dateofbirth)
	htmlString = strings.ReplaceAll(htmlString, "{{Nationality}}", user.Nationality)
	htmlString = strings.ReplaceAll(htmlString, "{{Summary}}", user.Summary)
	htmlString = strings.ReplaceAll(htmlString, "{{Workexperience}}", user.Workexperience)
	htmlString = strings.ReplaceAll(htmlString, "{{Education}}", user.Education)
	htmlString = strings.ReplaceAll(htmlString, "{{Skills}}", user.Skills)
	htmlString = strings.ReplaceAll(htmlString, "{{Languages}}", user.Languages)

	// Write 
	err = os.WriteFile("../bff/templates/populate_template.html", []byte(htmlString), 0644)
	if err != nil {
		panic(err)
	}
	// Read 
	populateHtml, err := os.ReadFile("../bff/templates/populate_template.html")
	if err != nil {
		log.Fatal(err)
	}
	// Create PDF 
	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}
	// Add HTML page 
	pdfg.AddPage(wkhtml.NewPageReader(bytes.NewReader(populateHtml)))
	// Create the PDF document in memory
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}
	// Write the PDF document to a file
	err = pdfg.WriteFile("./example.pdf")
	if err != nil {
		log.Fatal(err)
	}
	// Respond with template and user IDs
	fmt.Fprintf(w, "%s, %s", template_id, user_id)
}

	// * Get the value of an environment variable or return a default value
func getEnv(key, fallback string) string {
    if value, exists := os.LookupEnv(key); exists {
        return value
    }
    return fallback
}

 // ? Connection to the database
func connectToDatabase() (*sql.DB, error) {
	
    cfg := mysql.Config{
        User:                 getEnv("MYSQL_USER", os.Getenv("MYSQL_ROOT_USER")),
        Passwd:               getEnv("MYSQL_PASSWORD", os.Getenv("MYSQL_ROOT_PASSWORD")),
        Net:                  "tcp",
        Addr:                 os.Getenv("MYSQL_ADDR") + ":" + os.Getenv("MYSQL_PORT"),
        DBName:               os.Getenv("MYSQL_DATABASE"),
        AllowNativePasswords: true,
    }

	fmt.Println("\n * * * Establishing connection to the database...")
	fmt.Printf("\n\033[36m Environment variables printed fron main.go:\n\n")
    fmt.Printf("	User:          < %s >\n", cfg.User)
	fmt.Printf("	Password:      < %s*pass*%s >\n", string(cfg.Passwd[0]), string(cfg.Passwd[len(cfg.Passwd)-1]))
    fmt.Printf("	Address:       < %s >\n", cfg.Addr)
    fmt.Printf("	Database Name: < %s >\n\n", cfg.DBName)

    dsn := cfg.FormatDSN()
    maskedPasswd := fmt.Sprintf("%s*pass*%s", string(cfg.Passwd[0]), string(cfg.Passwd[len(cfg.Passwd)-1]))
    maskedDSN := fmt.Sprintf("%s:%s@tcp(%s)/%s?%s", cfg.User, maskedPasswd, cfg.Addr, cfg.DBName, dsn[strings.Index(dsn, "?")+1:])
    fmt.Printf(" DSN: %s\033[0m\n", maskedDSN)

    fmt.Println("\n * Opening database connection...")
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        fmt.Println("Error connection:", err)
        return nil, err
    }

    fmt.Println(" * Pinging DB...")
    if err = db.Ping(); err != nil {
        fmt.Printf("\033[31m	Error pinging database: %v\033[0m\n", err)
        db.Close()
        return nil, err
    }

    fmt.Println(" * Connected to database at the address:", cfg.Addr)
    return db, nil
}


func main() {
	var err error

	// Connect to the database
	db, err = connectToDatabase()
	if err != nil {
		log.Fatal(err) // Log and exit if there is an error
	}
	defer db.Close()

	// Create a new router
	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/users", homeUsers).Methods("GET")
	r.HandleFunc("/user", showUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/pdf", generateTemplate).Methods("GET")
	r.HandleFunc("/login", loginHandler).Methods("POST")
	r.HandleFunc("/signup", signupHandler).Methods("POST")
	r.HandleFunc("/logout", logoutHandler).Methods("POST")

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte("API is running"))
    }).Methods("GET")

	// Start the HTTP server
	log.Println("\n * Starting the HTTP server on port 8080...")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("\n * Failed to start HTTP server: %s\n", err)
	}
}
