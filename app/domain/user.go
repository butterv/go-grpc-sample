package domain

import (
	"database/sql"
)

//
type UserID string

//
type User struct {
	ID        UserID
	Email     string
	CreatedAt sql.NullTime
	UpdatedAt sql.NullTime
	DeletedAt sql.NullTime
}
