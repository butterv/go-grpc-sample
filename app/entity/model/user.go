package model

import (
	"time"
)

//
type UserID string

//
type User struct {
	ID        UserID
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}
