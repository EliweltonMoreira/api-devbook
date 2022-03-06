package controllers

import (
	"api/internal/db"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser insert a user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusUnprocessableEntity, err)
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	if err = user.Prepare(); err != nil {
		responses.Error(w, http.StatusBadRequest, err)
		return
	}

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	user.ID, err = repository.Create(user)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	responses.JSON(w, http.StatusCreated, user)
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
