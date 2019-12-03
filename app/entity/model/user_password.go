package model

import (
	"time"
)

//
type UserPassword struct {
	UserID       UserID
	PasswordHash string
	CreatedAt    time.Time
	UpdatedAt    time.Time
}
