package model

import (
	"time"
)

//
type UserTokenID int64

//
type UserToken struct {
	ID        UserTokenID
	UserID    UserID
	Token     string
	CreatedAt time.Time
}
