package repository

import (
	"github.com/istsh/go-grpc-sample/app/entity/model"
)

// UserTokenRepositoryAccess is a readonly repository for user tokens.
type UserTokenRepositoryAccess interface {
	FindByToken(token string) (*model.UserToken, error)
}

// UserTokenRepositoryModify is a write repository for user tokens.
type UserTokenRepositoryModify interface {
	UserTokenRepositoryAccess

	Create(userID model.UserID, token string) error
}
