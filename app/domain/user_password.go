package domain

import "database/sql"

//
type UserPassword struct {
	UserID       UserID
	PasswordHash string
	CreatedAt    sql.NullTime
	UpdatedAt    sql.NullTime
}
