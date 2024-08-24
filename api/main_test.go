// main_test.go
package main

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

// mock handlers
func mockHome(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Home"))
}

// mock users
func mockHomeUsers(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Users"))
}

// home handler
func TestHomeHandler(t *testing.T) {
    req, err := http.NewRequest("GET", "/", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    handler := http.HandlerFunc(mockHome)
    handler.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "Home"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

// router
func TestRouter(t *testing.T) {
    r := mux.NewRouter()
    r.HandleFunc("/", mockHome).Methods("GET")
    r.HandleFunc("/users", mockHomeUsers).Methods("GET")

    req, err := http.NewRequest("GET", "/users", nil)
    if err != nil {
        t.Fatal(err)
    }

    rr := httptest.NewRecorder()
    r.ServeHTTP(rr, req)

    if status := rr.Code; status != http.StatusOK {
        t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
    }

    expected := "Users"
    if rr.Body.String() != expected {
        t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
    }
}

// verify login
func TestVerifyLogin(t *testing.T) {
    var err error
    mockDB, mock, err := sqlmock.New()
    if err != nil {
        t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
    }
    defer mockDB.Close()

    // override db with mockDB
    originalDB := db
    db = mockDB
    defer func() { db = originalDB }()

    // generate hashed password
    password := "password123"
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    if err != nil {
        t.Fatalf("an error '%s' was not expected when hashing a password", err)
    }

    tests := []struct {
        name     string
        username string
        password string
        mock     func()
        want     bool
    }{
        {
            name:     "Valid login",
            username: "user1",
            password: "password123",
            mock: func() {
                mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
                    WithArgs("user1").
                    WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(string(hashedPassword)))
            },
            want: true,
        },
        {
            name:     "Invalid login",
            username: "user1",
            password: "wrongpassword",
            mock: func() {
                mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
                    WithArgs("user1").
                    WillReturnRows(sqlmock.NewRows([]string{"password"}).AddRow(string(hashedPassword)))
            },
            want: false,
        },
        {
            name:     "Database error",
            username: "user1",
            password: "password123",
            mock: func() {
                mock.ExpectQuery("SELECT password FROM userlogin WHERE username = ?").
                    WithArgs("user1").
                    WillReturnError(errors.New("db error"))
            },
            want: false,
        },
    }

    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            tt.mock()
            got := verifyLogin(tt.username, tt.password)
            if got != tt.want {
                t.Errorf("verifyLogin() = %v, want %v", got, tt.want)
            }
        })
    }
}
