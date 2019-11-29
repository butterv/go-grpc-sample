package persistence

import (
	"database/sql"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/infrastructure/repository/persistence/boil"
	"github.com/jinzhu/gorm"
)

type userRepository struct {
	db *sql.DB
}

func (r userRepository) Find(id model.UserID) (*model.User, error) {
	u := boil.User{
		ID: string(id),
	}
	u.

	user := &model.User{}
	if err := r.connection.Where("id = ?", id).First(user).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
