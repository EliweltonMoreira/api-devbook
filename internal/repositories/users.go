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

// Update change a user info in the database
func (repository Users) Update(ID uint64, user models.User) error {
	statement, err := repository.db.Prepare(
		"update users set name = $1, nick = $2, email = $3 where id = $4",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(user.Name, user.Nick, user.Email, ID); err != nil {
		return err
	}

	return nil
}

// Delete remove a user info in the database
func (repository Users) Delete(ID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from users where id = $1",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(ID); err != nil {
		return err
	}

	return nil
}

// GetByEmail get a user by email and returns their id and hashed password
func (repository Users) GetByEmail(email string) (models.User, error) {
	line, err := repository.db.Query("select id, password from users where email = $1", email)
	if err != nil {
		return models.User{}, err
	}
	defer line.Close()

	var user models.User
	if line.Next() {
		if err = line.Scan(&user.ID, &user.Password); err != nil {
			return models.User{}, err
		}
	}

	return user, nil
}

// Follow allows a user follow other user
func (repository Users) Follow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"insert into followers (user_id, follower_id) values ($1, $2) on conflict (user_id, follower_id) do nothing",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// Unfollow allows a user unfollow other user
func (repository Users) Unfollow(userID, followerID uint64) error {
	statement, err := repository.db.Prepare(
		"delete from followers where user_id = $1 and follower_id = $2",
	)
	if err != nil {
		return err
	}
	defer statement.Close()

	if _, err = statement.Exec(userID, followerID); err != nil {
		return err
	}

	return nil
}

// GetFollowers get all followers of a user
func (repository Users) GetFollowers(userID uint64) ([]models.User, error) {
	lines, err := repository.db.Query(`
		select u.id, u.name, u.nick, u.email, u.created_at
		from users u
		inner join followers f on u.id = f.follower_id
		where f.user_id = $1`,
		userID,
	)
	if err != nil {
		return nil, err
	}
	defer lines.Close()

	var users []models.User
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
