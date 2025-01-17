// Code generated by sqlc. DO NOT EDIT.

package db

import (
	"time"
)

type User struct {
	Username         string    `json:"username"`
	HashedPassword   string    `json:"hashed_password"`
	FullName         string    `json:"full_name"`
	Email            string    `json:"email"`
	PasswordChangeAt time.Time `json:"password_change_at"`
	CreatedAt        time.Time `json:"created_at"`
}
