package repository

import (
	"github.com/istsh/go-grpc-sample/app/entity/model"
)

// UserPasswordRepositoryAccess is a readonly repository for user passwords.
type UserPasswordRepositoryAccess interface {
	Find(userID model.UserID) (*model.UserPassword, error)
}

// UserPasswordRepositoryModify is a read/write repository for user passwords.
type UserPasswordRepositoryModify interface {
	UserPasswordRepositoryAccess

	Create(userID model.UserID, passwordHash string) error
	Update(userID model.UserID, passwordHash string) error
}
