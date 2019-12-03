package persistence

import (
	"github.com/jinzhu/gorm"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type dbRepository struct {
	db *gorm.DB
}

type dbConnection struct {
	db *gorm.DB
}

type dbTransaction struct {
	db *gorm.DB
}

// NewDBRepository generates a new repository using DB.
func NewDBRepository(db *gorm.DB) repository.Repository {
	return &dbRepository{
		db: db,
	}
}

func (r dbRepository) NewConnection() repository.Connection {
	return &dbConnection{
		db: r.db,
	}
}

func (con *dbConnection) Close() error {
	// We don't need to close *sql.DB. No need to do anything.
	return nil
}

func (con *dbConnection) RunTransaction(f func(repository.Transaction) error) error {
	tx := con.db.Begin()

	err := f(&dbTransaction{db: tx})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit().Error
	if err != nil {
		return err
	}

	return nil
}

func (con *dbConnection) User() repository.UserRepositoryAccess {
	return dbUserRepository{db: con.db}
}

func (con *dbConnection) UserPassword() repository.UserPasswordRepositoryAccess {
	return dbUserPasswordRepository{db: con.db}
}

func (con *dbConnection) UserToken() repository.UserTokenRepositoryAccess {
	return dbUserTokenRepository{db: con.db}
}

func (tx *dbTransaction) User() repository.UserRepositoryModify {
	return dbUserRepository{db: tx.db}
}

func (tx *dbTransaction) UserPassword() repository.UserPasswordRepositoryModify {
	return dbUserPasswordRepository{db: tx.db}
}

func (tx *dbTransaction) UserToken() repository.UserTokenRepositoryModify {
	return dbUserTokenRepository{db: tx.db}
}
