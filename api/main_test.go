// main_test.go
package main

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
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

// set session
func TestSetSession(t *testing.T) {
	userName := "testUser"
	recorder := httptest.NewRecorder()

	setSession(userName, recorder)

	result := recorder.Result()
	defer result.Body.Close()

	if result.StatusCode != http.StatusOK {
		t.Errorf("expected status OK; got %v", result.Status)
	}

}

// logout handler
func TestLogoutHandler_ClearsSession(t *testing.T) {
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(logoutHandler)

	handler.ServeHTTP(rr, req)

	cookie := rr.Result().Cookies()
	if len(cookie) == 0 || cookie[0].Value != "" || cookie[0].MaxAge != -1 {
		t.Errorf("Expected session cookie to be cleared, got %v", cookie)
	}
}

func TestLogoutHandler_StatusNoContent(t *testing.T) {
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(logoutHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusNoContent {
		t.Errorf("Expected status code %v, got %v", http.StatusNoContent, status)
	}
}

func TestLogoutHandler_ResponseMessage(t *testing.T) {
	req, err := http.NewRequest("GET", "/logout", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(logoutHandler)

	handler.ServeHTTP(rr, req)

	expected := "User logout successfully"
	if rr.Body.String() != expected {
		t.Errorf("Expected response body %v, got %v", expected, rr.Body.String())
	}
}

// signup handler
func TestSignupHandler_ParseFormError(t *testing.T) {
	req, err := http.NewRequest("POST", "/signup", strings.NewReader("invalid=form"))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(signupHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, status)
	}
}

func TestSignupHandler_EmptyEmailOrPassword(t *testing.T) {
	req, err := http.NewRequest("POST", "/signup", strings.NewReader("email=&password="))
	if err != nil {
		t.Fatal(err)
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(signupHandler)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("Expected status code %v, got %v", http.StatusBadRequest, status)
	}
}

// home users
func TestHomeUsers(t *testing.T) {
	var mock sqlmock.Sqlmock
	db, mock, _ = sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "firstname", "lastname", "email"}).
		AddRow(1, "John", "Doe", "john.doe@example.com").
		AddRow(2, "Jane", "Doe", "jane.doe@example.com")

	mock.ExpectQuery("SELECT id, firstname, lastname, email FROM users").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/users", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(homeUsers)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var gotUsers []User
	if err := json.Unmarshal(rr.Body.Bytes(), &gotUsers); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expectedUsers := []User{
		{ID: 1, Firstname: "John", Lastname: "Doe", Email: "john.doe@example.com"},
		{ID: 2, Firstname: "Jane", Lastname: "Doe", Email: "jane.doe@example.com"},
	}

	if !reflect.DeepEqual(gotUsers, expectedUsers) {
		t.Errorf("handler returned unexpected body: got %v want %v", gotUsers, expectedUsers)
	}
}

// home handler
func TestHomeHandler_NonGETMethod(t *testing.T) {
	req, err := http.NewRequest("POST", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusMethodNotAllowed {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusMethodNotAllowed)
	}

	expected := "Method Not Allowed"
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", rr.Body.String(), expected)
	}
}

func TestHomeHandler_Success(t *testing.T) {
	var mock sqlmock.Sqlmock
	db, mock, _ = sqlmock.New()
	defer db.Close()

	rows := sqlmock.NewRows([]string{"id", "jobtitle", "firstname", "lastname", "email", "phone", "address", "city", "country", "postalcode", "dateofbirth", "nationality", "summary", "workexperience", "education", "skills", "languages"}).
		AddRow(1, "Developer", "John", "Doe", "john.doe@example.com", "1234567890", "123 Street", "City", "Country", "12345", "1990-01-01", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages").
		AddRow(2, "Manager", "Jane", "Doe", "jane.doe@example.com", "0987654321", "456 Avenue", "City", "Country", "67890", "1985-01-01", "Nationality", "Summary", "Work Experience", "Education", "Skills", "Languages")

	mock.ExpectQuery("SELECT id, jobtitle, firstname, lastname, email, phone, address, city, country, postalcode, dateofbirth, nationality, summary, workexperience, education, skills, languages FROM users").WillReturnRows(rows)

	req, err := http.NewRequest("GET", "/home", nil)
	if err != nil {
		t.Fatal(err)
	}

	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(home)

	handler.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}

	var gotUsers []User
	if err := json.Unmarshal(rr.Body.Bytes(), &gotUsers); err != nil {
		t.Fatalf("could not unmarshal response: %v", err)
	}

	expectedUsers := []User{
		{ID: 1, Jobtitle: "Developer", Firstname: "John", Lastname: "Doe", Email: "john.doe@example.com", Phone: "1234567890", Address: "123 Street", City: "City", Country: "Country", Postalcode: "12345", Dateofbirth: "1990-01-01", Nationality: "Nationality", Summary: "Summary", Workexperience: "Work Experience", Education: "Education", Skills: "Skills", Languages: "Languages"},
		{ID: 2, Jobtitle: "Manager", Firstname: "Jane", Lastname: "Doe", Email: "jane.doe@example.com", Phone: "0987654321", Address: "456 Avenue", City: "City", Country: "Country", Postalcode: "67890", Dateofbirth: "1985-01-01", Nationality: "Nationality", Summary: "Summary", Workexperience: "Work Experience", Education: "Education", Skills: "Skills", Languages: "Languages"},
	}

	if !reflect.DeepEqual(gotUsers, expectedUsers) {
		t.Errorf("handler returned unexpected body: got %v want %v", gotUsers, expectedUsers)
	}
}
