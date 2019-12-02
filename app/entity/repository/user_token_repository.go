package repository

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

// UserTokenRepositoryAccess is a readonly repository for user tokens.
type UserTokenRepositoryAccess interface {
	FindByToken(ctx context.Context, token string) (*model.UserToken, error)
}

// UserTokenRepositoryModify is a write repository for user tokens.
type UserTokenRepositoryModify interface {
	Create(ctx context.Context, userID model.UserID, token string) error
	Delete(ctx context.Context, id model.UserTokenID) error
}
