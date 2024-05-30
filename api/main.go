package main

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

func home(w http.ResponseWriter, r *http.Request) {
<<<<<<< HEAD:project.go
=======

>>>>>>> 1a8974f385c67f1a1622514aaa89ee36ab720afb:api/main.go
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

<<<<<<< HEAD:project.go
	var users []User
=======
>>>>>>> 1a8974f385c67f1a1622514aaa89ee36ab720afb:api/main.go
	rows, err := db.Query("SELECT id, jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

<<<<<<< HEAD:project.go
=======
	var users []User

>>>>>>> 1a8974f385c67f1a1622514aaa89ee36ab720afb:api/main.go
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
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", 1)
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
<<<<<<< HEAD:project.go
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE users SET jobtitle = ?, firstname = ?, lastname = ?, email = ?, phone = ?, address = ?, city = ?, country = ?, postalcode = ?, dateofbirth = ?, nationality = ?, summary = ?, workexperience = ?, education = ?, skills = ?, languages = ? WHERE id = ?", 
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	user.ID = int64(id)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)

=======
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	_, err = db.Exec("UPDATE users SET jobtitle=?, first_name=?, last_name=?, email=?, phone=?, address=?, city=?, country=?, postal_code=?, date_of_birth=?, nationality=?, summary=?, work_experience=?, education=?, skills=?, languages=? WHERE id=?",
		user.Jobtitle, user.Firstname, user.Lastname, user.Email, user.Phone, user.Address, user.City, user.Country, user.Postalcode, user.Dateofbirth, user.Nationality, user.Summary, user.Workexperience, user.Education, user.Skills, user.Languages, id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	user.ID = int64(id)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
>>>>>>> 1a8974f385c67f1a1622514aaa89ee36ab720afb:api/main.go
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
		fmt.Fprintf(w, "ID should be a integer")
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
		fmt.Println(err)
	}

	idtemplate_int, err := strconv.Atoi(template_id[0])
	if err != nil {
		fmt.Println(err)
	}

	var user User
	var template Template

	row1 := db.QueryRow("SELECT Path FROM template WHERE id = ?", idtemplate_int)
	row1.Scan(&template.Path)
	row := db.QueryRow("SELECT * FROM user WHERE id = ?", iduser_int)
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

	err = os.WriteFile("../bff/templates/populate_template.html", []byte(htmlString), 0644)
	if err != nil {
		panic(err)
	}

	populateHtml, err := os.ReadFile("../bff/templates/populate_template.html")
	if err != nil {
		log.Fatal(err)
	}

	pdfg, err := wkhtml.NewPDFGenerator()
	if err != nil {
		return
	}

	pdfg.AddPage(wkhtml.NewPageReader(bytes.NewReader(populateHtml)))

	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile("./example.pdf")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(w, "%s, %s", template_id, user_id)
}

func connectToDatabase() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"), // TODO
		Passwd:               os.Getenv("DBPASS"),
		Net:                  "tcp",
		Addr:                 "127.0.0.1:3306",
		DBName:               "users",
		AllowNativePasswords: true,
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, err
	}

	fmt.Println("Connected to database!")
	return db, nil
}

func main() {
	var err error
	db, err = connectToDatabase()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := mux.NewRouter()

	r.HandleFunc("/", home).Methods("GET")
	r.HandleFunc("/user", showUser).Methods("GET")
	r.HandleFunc("/user", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	r.HandleFunc("/pdf", generateTemplate).Methods("GET")

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("listen and serve: %s\n", err)
	}
}
