package main

import (
	"github.com/gorilla/mux"
	"go-mongo-crud/user"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	// Lists Users
	r.HandleFunc("/api/users/", user.GetAllUsers).Methods("GET")
	// Get User By Id
	r.HandleFunc("/api/users/{id}", user.GetUser).Methods("GET")
	// Create User
	r.HandleFunc("/api/users/", user.CreateUser).Methods("POST")
	// Delete User by Id
	r.HandleFunc("/api/users/{id}", user.DeleteUser).Methods("DELETE")
	// Update User by Id
	r.HandleFunc("/api/users/{id}", user.UpdateUser).Methods("PUT")
	http.ListenAndServe(":3000", r)
}
