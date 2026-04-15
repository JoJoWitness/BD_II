package models

import (
	"database/sql"
	"errors"

	"demo/config"
)

type User struct {
	ID       int
	Username string
	Password string
}

var ErrUserNotFound = errors.New("user not found")

func FindUserByUsername(username string) (*User, error) {
	var u User
	err := config.DB.QueryRow(
		`SELECT id, username, password FROM users WHERE username=$1`,
		username,
	).Scan(&u.ID, &u.Username, &u.Password)
	if err == sql.ErrNoRows {
		return nil, ErrUserNotFound
	}
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func CreateUser(username, password string) (*User, error) {
	u := &User{Username: username, Password: password}
	err := config.DB.QueryRow(
		`INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id`,
		username, password,
	).Scan(&u.ID)
	if err != nil {
		return nil, err
	}
	return u, nil
}
