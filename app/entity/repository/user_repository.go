package repository

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

// UserRepositoryAccess is a readonly repository for users.
type UserRepositoryAccess interface {
	Find(ctx context.Context, id model.UserID) (*model.User, error)
	FindByEmail(ctx context.Context, email string) (*model.User, error)
}

// UserRepositoryModify is a write repository for users.
type UserRepositoryModify interface {
	Create(ctx context.Context, id model.UserID, email string) error
}
