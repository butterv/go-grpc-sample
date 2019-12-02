package persistence

import (
	"context"
	"database/sql"
	"errors"

	sb "github.com/volatiletech/sqlboiler/boil"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/infrastructure/repository/persistence/boil"
)

type dbUserPasswordRepository struct {
	db *sql.DB
}

type txUserPasswordRepository struct {
	tx *sql.Tx
}

func toUserPasswordModel(up *boil.UserPassword) *model.UserPassword {
	return &model.UserPassword{
		UserID:       model.UserID(up.UserID),
		PasswordHash: up.PasswordHash,
		CreatedAt: sql.NullTime{
			Time:  up.CreatedAt,
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  up.UpdatedAt,
			Valid: true,
		},
	}
}

func (r dbUserPasswordRepository) Find(ctx context.Context, userID model.UserID) (*model.UserPassword, error) {
	qm := boil.UserPasswordWhere.UserID.EQ(string(userID))

	up, err := boil.UserPasswords(qm).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return toUserPasswordModel(up), nil
}

func (r txUserPasswordRepository) Create(ctx context.Context, userID model.UserID, passwordHash string) error {
	up := boil.UserPassword{
		UserID:       string(userID),
		PasswordHash: passwordHash,
	}
	return up.Insert(ctx, r.tx, sb.Infer())
}

func (r txUserPasswordRepository) Update(ctx context.Context, userID model.UserID, passwordHash string) error {
	up := boil.UserPassword{
		UserID:       string(userID),
		PasswordHash: passwordHash,
	}
	_, err := up.Update(ctx, r.tx, sb.Infer())
	return err
}
