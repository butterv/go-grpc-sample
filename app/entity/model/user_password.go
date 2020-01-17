package model

import (
	"time"
)

// UserPassword has a user id and password hash.
type UserPassword struct {
	UserID       UserID
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
