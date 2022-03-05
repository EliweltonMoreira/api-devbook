package controllers

import "net/http"

// CreateUser insert a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user!"))
}

// GetUsers get all users stored in the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users!"))
}

// GetUser get one user stored in the database
func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting a user!"))
}

// UpdateUser change a user info in the database
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating user!"))
}

// DeleteUser remove a user info in the database
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Removing user!"))
}
