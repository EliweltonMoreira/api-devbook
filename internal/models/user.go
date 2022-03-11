package models

import (
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// User represents a user using the social network
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nick      string    `json:"nick,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// Prepare will call the methods to validate and format the received user
func (user *User) Prepare(stage string) error {
	if err := user.validate(stage); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate(stage string) error {
	if user.Name == "" {
		return errors.New("the name is required and cannot be blank")
	}

	if user.Nick == "" {
		return errors.New("the nick is required and cannot be blank")
	}

	if user.Email == "" {
		return errors.New("the e-mail is required and cannot be blank")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("the inserted e-mail is invalid")
	}

	if stage == "register" && user.Password == "" {
		return errors.New("the password is required and cannot be blank")
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)
}
