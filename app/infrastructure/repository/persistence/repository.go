package persistence

import (
	"database/sql"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type dbRepository struct {
	db *sql.DB
}

type dbConnection struct {
	db *sql.DB
}

type dbTransaction struct {
	tx *sql.Tx
}

// NewDBRepository generates a new repository using DB.
func NewDBRepository(db *sql.DB) repository.Repository {
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
	tx, err := con.db.Begin()
	if err != nil {
		return err
	}

	err = f(&dbTransaction{tx: tx})
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
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
	return txUserRepository{tx: tx.tx}
}

func (tx *dbTransaction) UserPassword() repository.UserPasswordRepositoryModify {
	return txUserPasswordRepository{tx: tx.tx}
}

func (tx *dbTransaction) UserToken() repository.UserTokenRepositoryModify {
	return txUserTokenRepository{tx: tx.tx}
}
