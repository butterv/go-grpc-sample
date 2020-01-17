package model

import (
	"time"
)

// UserTokenID is an id type for user token.
type UserTokenID int64

// UserToken has a id, user id and token.
type UserToken struct {
	ID        UserTokenID
	UserID    UserID
	Token     string
	CreatedAt time.Time
}
