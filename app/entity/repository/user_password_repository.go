package repository

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

// UserPasswordRepositoryAccess is a readonly repository for user passwords.
type UserPasswordRepositoryAccess interface {
	Find(ctx context.Context, userID model.UserID) (*model.UserPassword, error)
}

// UserPasswordRepositoryModify is a write repository for user passwords.
type UserPasswordRepositoryModify interface {
	Create(ctx context.Context, userID model.UserID, passwordHash string) error
	Update(ctx context.Context, userID model.UserID, passwordHash string) error
}
