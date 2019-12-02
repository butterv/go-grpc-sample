package persistence

import (
	"context"
	"database/sql"
	"errors"

	sb "github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/infrastructure/repository/persistence/boil"
)

type dbUserRepository struct {
	db *sql.DB
}

type txUserRepository struct {
	tx *sql.Tx
}

func toUserModel(u *boil.User) *model.User {
	return &model.User{
		ID:    model.UserID(u.ID),
		Email: u.Email,
		CreatedAt: sql.NullTime{
			Time:  u.CreatedAt,
			Valid: true,
		},
		UpdatedAt: sql.NullTime{
			Time:  u.UpdatedAt,
			Valid: true,
		},
		DeletedAt: sql.NullTime{
			Time:  u.DeletedAt.Time,
			Valid: u.DeletedAt.Valid,
		},
	}
}

func (r dbUserRepository) Find(ctx context.Context, id model.UserID) (*model.User, error) {
	qms := []qm.QueryMod{
		boil.UserWhere.ID.EQ(string(id)),
		boil.UserWhere.DeletedAt.IsNull(),
	}

	u, err := boil.Users(qms...).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return toUserModel(u), nil
}

func (r dbUserRepository) FindByEmail(ctx context.Context, email string) (*model.User, error) {
	qms := []qm.QueryMod{
		boil.UserWhere.Email.EQ(email),
		boil.UserWhere.DeletedAt.IsNull(),
	}

	u, err := boil.Users(qms...).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return toUserModel(u), nil
}

func (r txUserRepository) Create(ctx context.Context, id model.UserID, email string) error {
	u := boil.User{
		ID:    string(id),
		Email: email,
	}
	return u.Insert(ctx, r.tx, sb.Infer())
}
