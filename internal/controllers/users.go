package controllers

import (
	"api/internal/db"
	"api/internal/models"
	"api/internal/repositories"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser insert a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	repository := repositories.NewRepositoryOfUsers(db)
	userID, err := repository.Create(user)
	if err != nil {
		log.Fatal(err)
	}

	w.Write([]byte(fmt.Sprintf("Inserted ID: %d", userID)))
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
