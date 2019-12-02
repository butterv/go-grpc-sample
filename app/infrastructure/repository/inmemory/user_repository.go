package inmemory

import (
	"context"
	"database/sql"
	"time"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type inmemoryUserRepository struct {
	repository.UserRepositoryAccess

	s *Store
}

type inmemoryTxUserRepository struct {
	repository.UserRepositoryModify

	s *Store
}

func (r inmemoryUserRepository) Find(ctx context.Context, id model.UserID) (*model.User, error) {
	for _, u := range r.s.users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, nil
}

func (r inmemoryUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	for _, u := range r.s.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, nil
}

func (r inmemoryTxUserRepository) Create(ctx context.Context, id model.UserID, email string) error {
	now := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	u := &model.User{
		ID:        id,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
	r.s.AddUsers(u)

	return nil
}
