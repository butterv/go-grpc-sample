package inmemory

import (
	"time"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type inmemoryUserRepository struct {
	repository.UserRepositoryAccess
	repository.UserRepositoryModify

	s *Store
}

func (r inmemoryUserRepository) Find(id model.UserID) (*model.User, error) {
	for _, u := range r.s.users {
		if u.ID == id {
			return u, nil
		}
	}

	return nil, nil
}

func (r inmemoryUserRepository) FindByEmail(email string) (*model.User, error) {
	for _, u := range r.s.users {
		if u.Email == email {
			return u, nil
		}
	}

	return nil, nil
}

func (r inmemoryUserRepository) Create(id model.UserID, email string) error {
	now := time.Now()

	u := &model.User{
		ID:        id,
		Email:     email,
		CreatedAt: now,
		UpdatedAt: now,
	}
	r.s.AddUsers(u)

	return nil
}
