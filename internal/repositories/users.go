package repositories

import (
	"api/internal/models"
	"database/sql"
	"fmt"
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

// Get brings up all users that match a name or nick filter
func (repository Users) Get(nameOrNick string) ([]models.User, error) {
	nameOrNick = fmt.Sprintf("%%%s%%", nameOrNick)

	lines, err := repository.db.Query(
		"select id, name, nick, email, created_at from users where name LIKE $1 or nick LIKE $2",
		nameOrNick, nameOrNick,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	users := []models.User{}
	for lines.Next() {
		var user models.User

		if err = lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	return users, nil
}

// GetByID brings up a user in the database
func (repository Users) GetByID(ID uint64) (models.User, error) {
	lines, err := repository.db.Query(
		"select id, name, nick, email, created_at from users where id = $1",
		ID,
	)
	if err != nil {
		return models.User{}, err
	}
	defer lines.Close()

	var user models.User
	if lines.Next() {
		if err := lines.Scan(
			&user.ID,
			&user.Name,
			&user.Nick,
			&user.Email,
			&user.CreatedAt,
		); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}
