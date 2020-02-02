package user

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)


// Handle Error
func handleError(err error, message string, w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	w.Write([]byte(fmt.Sprintf(message, err)))
}

// GetAllUsers returns a list of all database items to the response.
func GetAllUsers(w http.ResponseWriter, req *http.Request) {
	rs, err := GetAll()
	if err != nil {
		handleError(err, "Failed to load users: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to load marshal data: %v", w)
		return
	}

	w.Write(bs)
}

// GetUser returns a single database item matching given ID parameter.
func GetUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	value := vars["id"]
	rs, err := GetOne(value)
	if err != nil {
		handleError(err, "Failed to read user: %v", w)
		return
	}

	bs, err := json.Marshal(rs)
	if err != nil {
		handleError(err, "Failed to marshal data: %v", w)
		return
	}

	w.Write(bs)
}

// CreateUser saves an user (form data) into the database.
func CreateUser(w http.ResponseWriter, req *http.Request) {
	ID := req.FormValue("id")
	firstNameStr := req.FormValue("firstname")
	lastNameStr := req.FormValue("lastname")
	emailStr := req.FormValue("email")

	user := User{
		ID: ID, FirstName: firstNameStr, LastName: lastNameStr, EmailID: emailStr,
	}

	if err := Save(user); err != nil {
		handleError(err, "Failed to save user: %v", w)
		return
	}

	w.Write([]byte("OK"))
}

// DeleteUser removes a single user (identified by parameter) from the database.
func DeleteUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]

	if err := Remove(id); err != nil {
		handleError(err, "Failed to remove user: %v", w)
		return
	}

	w.Write([]byte("OK"))
}


// UpdateUser update a user(form data) from the database.
func UpdateUser(w http.ResponseWriter, req *http.Request) {
	vars := mux.Vars(req)
	id := vars["id"]
	firstNameStr := req.FormValue("firstname")
	lastNameStr := req.FormValue("lastname")
	emailStr := req.FormValue("email")
	user := User{
		ID: id, FirstName: firstNameStr, LastName: lastNameStr, EmailID: emailStr,
	}
	err := Update(user)
	if err != nil{
		handleError(err, "Unable to update user", w)
	}
	w.Write([]byte("OK"))
}
