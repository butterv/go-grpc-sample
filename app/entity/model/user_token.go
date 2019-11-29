package model

import "database/sql"

//
type UserTokenID int64

//
type UserToken struct {
	ID        UserTokenID
	UserID    UserID
	Token     string
	CreatedAt sql.NullTime
	DeletedAt sql.NullTime
}
