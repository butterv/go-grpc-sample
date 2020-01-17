package model

import (
	"time"
)

// UserID is an id type for user.
type UserID string

// User represents a user.
type User struct {
	ID        UserID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
