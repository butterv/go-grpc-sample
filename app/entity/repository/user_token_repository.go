package repository

import "github.com/istsh/go-grpc-sample/app/entity/model"

// UserTokenRepositoryAccess is a readonly repository for user tokens.
type UserTokenRepositoryAccess interface {
	FindByToken(token string) (*model.UserToken, error)
}

// UserTokenRepositoryModify is a read/write repository for user tokens.
type UserTokenRepositoryModify interface {
	UserTokenRepositoryAccess

	Create(userID model.UserID, passwordHash string) error
	Update(userID model.UserID, passwordHash string) error
}
