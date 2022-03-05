package repositories

import (
	"api/internal/models"
	"database/sql"
)

type Users struct {
	db *sql.DB
}

// NewRepositoryOfUsers create a repository of users
func NewRepositoryOfUsers(db *sql.DB) *Users {
	return &Users{db}
}

// Create insert a user in the database
func (repository Users) Create(user models.User) (uint64, error) {
	statement, err := repository.db.Prepare(
		"insert into users (name, nick, email, password) values ($1, $2, $3, $4) returning id",
	)
	if err != nil {
		return 0, err
	}
	defer statement.Close()

	var userID int64
	err = statement.QueryRow(user.Name, user.Nick, user.Email, user.Password).Scan(&userID)
	if err != nil {
		return 0, err
	}

	return uint64(userID), nil
}
