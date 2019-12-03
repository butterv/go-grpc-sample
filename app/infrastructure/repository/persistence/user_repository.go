package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

type dbUserRepository struct {
	db *gorm.DB
}

func (r dbUserRepository) Find(id model.UserID) (*model.User, error) {
	user := &model.User{}
	if err := r.db.Where("id = ?", id).First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}

func (r dbUserRepository) FindByEmail(email string) (*model.User, error) {
	u := &model.User{}
	if err := r.db.Where("email = ?", email).First(u).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return u, nil
}

func (r dbUserRepository) Create(id model.UserID, email string) error {
	u := &model.User{
		ID:    id,
		Email: email,
	}
	return r.db.Create(u).Error
}
