package inmemory

import (
	"time"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type inmemoryUserPasswordRepository struct {
	repository.UserPasswordRepositoryAccess
	repository.UserPasswordRepositoryModify

	s *Store
}

func (r inmemoryUserPasswordRepository) Find(userID model.UserID) (*model.UserPassword, error) {
	for _, up := range r.s.userPasswords {
		if up.UserID == userID {
			return up, nil
		}
	}

	return nil, nil
}

func (r inmemoryUserPasswordRepository) Create(userID model.UserID, passwordHash string) error {
	now := time.Now()

	up := &model.UserPassword{
		UserID:       userID,
		PasswordHash: passwordHash,
		CreatedAt:    now,
		UpdatedAt:    now,
	}
	r.s.AddUserPasswords(up)

	return nil
}

func (r inmemoryUserPasswordRepository) Update(userID model.UserID, passwordHash string) error {
	for i, up := range r.s.userPasswords {
		if up.UserID == userID {
			now := time.Now()
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
