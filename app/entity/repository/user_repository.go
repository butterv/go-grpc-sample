package repository

import (
	"github.com/istsh/go-grpc-sample/app/entity/model"
)

// UserRepositoryAccess is a readonly repository for users.
type UserRepositoryAccess interface {
	Find(id model.UserID) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
}

// UserRepositoryModify is a read/write repository for users.
type UserRepositoryModify interface {
	UserRepositoryAccess

	Create(id model.UserID, email string) error
}
