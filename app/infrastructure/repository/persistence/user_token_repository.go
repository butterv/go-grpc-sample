package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

type dbUserTokenRepository struct {
	db *gorm.DB
}

func (r dbUserTokenRepository) FindByToken(token string) (*model.UserToken, error) {
	ut := &model.UserToken{}
	if err := r.db.Where("token = ?", token).First(ut).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return ut, nil
}

func (r dbUserTokenRepository) Create(userID model.UserID, token string) error {
	ut := &model.UserToken{
		UserID: userID,
		Token:  token,
	}
	return r.db.Create(ut).Error
}
