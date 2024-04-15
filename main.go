package main

import (
	"encoding/json"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
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

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	w.Write([]byte("Welcome!"))
}

func showUser(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		w.WriteHeader(405)
		w.Write([]byte("Method Not Allowed"))
		return
	}

	id, err := strconv.Atoi(r.URL.Query().Get("id")) // to convert the string value to an integer
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	var user User
    row := db.QueryRow("SELECT * FROM user WHERE id = ?", id)
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

	fmt.Fprintf(w, "Update the user with ID %d", id)
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

func connectToDatabase() (*sql.DB, error) {
	cfg := mysql.Config{
		User:                 os.Getenv("DBUSER"),
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
	r.HandleFunc("/user/{id}", showUser).Methods("POST")
	r.HandleFunc("/user/", createUser).Methods("POST")
	r.HandleFunc("/user/{id}", updateUser).Methods("PUT")
	r.HandleFunc("/user/{id}", deleteUser).Methods("DELETE")
	http.Handle("/", r)

	log.Println("Starting server on :8080")

	if err := http.ListenAndServe("8080", r); err != nil {
		log.Fatalf("listen and serve: %s\n", err)
	}
}
