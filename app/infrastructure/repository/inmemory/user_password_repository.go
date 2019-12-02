package inmemory

import (
	"context"
	"database/sql"
	"time"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type inmemoryUserPasswordRepository struct {
	repository.UserPasswordRepositoryAccess

	s *Store
}

type inmemoryTxUserPasswordRepository struct {
	repository.UserPasswordRepositoryModify

	s *Store
}

func (r inmemoryUserPasswordRepository) Find(ctx context.Context, userID model.UserID) (*model.UserPassword, error) {
	for _, up := range r.s.userPasswords {
		if up.UserID == userID {
			return up, nil
		}
	}

	return nil, nil
}

func (r inmemoryTxUserPasswordRepository) Create(ctx context.Context, userID model.UserID, passwordHash string) error {
	now := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	up := &model.UserPassword{
		UserID:       userID,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	r.s.AddUserPasswords(up)

	return nil
}

func (r inmemoryTxUserPasswordRepository) Update(ctx context.Context, userID model.UserID, passwordHash string) error {
	for i, up := range r.s.userPasswords {
		if up.UserID == userID {
			now := sql.NullTime{
				Time:  time.Now(),
				Valid: true,
			}
			r.s.userPasswords[i] = &model.UserPassword{
				UserID:       userID,
				PasswordHash: passwordHash,
				CreatedAt:    now,
				UpdatedAt:    now,
			}
			break
		}
	}

	return nil
}
