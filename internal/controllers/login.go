package controllers

import (
	"api/internal/authentication"
	"api/internal/db"
	"api/internal/models"
	"api/internal/repositories"
	"api/internal/responses"
	"api/internal/security"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Login is responsible for authenticating a user in the api
func Login(w http.ResponseWriter, r *http.Request) {
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

	db, err := db.Connect()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()

	repository := repositories.NewRepositoryOfUsers(db)
	userSavedInTheDB, err := repository.GetByEmail(user.Email)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, err)
		return
	}

	if err = security.VerifyPassword(userSavedInTheDB.Password, user.Password); err != nil {
		responses.Error(w, http.StatusUnauthorized, err)
		return
	}

	token, _ := authentication.CreateToken(userSavedInTheDB.ID)
	fmt.Println(token)
	w.Write([]byte(token))
}
