package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

type dbUserPasswordRepository struct {
	db *gorm.DB
}

func (r dbUserPasswordRepository) Find(userID model.UserID) (*model.UserPassword, error) {
	up := &model.UserPassword{}
	if err := r.db.Where("user_id = ?", userID).First(up).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return up, nil
}

func (r dbUserPasswordRepository) Create(userID model.UserID, passwordHash string) error {
	up := &model.UserPassword{
		UserID:       userID,
		PasswordHash: passwordHash,
	}
	return r.db.Create(up).Error
}

func (r dbUserPasswordRepository) Update(userID model.UserID, passwordHash string) error {
	userPassword := &model.UserPassword{
		UserID: userID,
	}
	return r.db.Model(userPassword).Update("password_hash", passwordHash).Error
}
