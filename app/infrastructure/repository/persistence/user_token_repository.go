package persistence

import (
	"context"
	"database/sql"
	"errors"
	"time"

	"github.com/volatiletech/null"
	sb "github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries/qm"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/infrastructure/repository/persistence/boil"
)

type dbUserTokenRepository struct {
	db *sql.DB
}

type txUserTokenRepository struct {
	tx *sql.Tx
}

func toUserTokenModel(ut *boil.UserToken) *model.UserToken {
	return &model.UserToken{
		ID:     model.UserTokenID(ut.ID),
		UserID: model.UserID(ut.UserID),
		Token:  ut.Token,
		CreatedAt: sql.NullTime{
			Time:  ut.CreatedAt,
			Valid: true,
		},
		DeletedAt: sql.NullTime{
			Time:  ut.DeletedAt.Time,
			Valid: ut.DeletedAt.Valid,
		},
	}
}

func (r dbUserTokenRepository) FindByToken(ctx context.Context, token string) (*model.UserToken, error) {
	qms := []qm.QueryMod{
		boil.UserTokenWhere.Token.EQ(token),
		boil.UserTokenWhere.DeletedAt.IsNull(),
	}

	ut, err := boil.UserTokens(qms...).One(ctx, r.db)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return toUserTokenModel(ut), nil
}

func (r txUserTokenRepository) Create(ctx context.Context, userID model.UserID, token string) error {
	ut := boil.UserToken{
		UserID: string(userID),
		Token:  token,
	}
	return ut.Insert(ctx, r.tx, sb.Infer())
}

func (r txUserTokenRepository) Delete(ctx context.Context, id model.UserTokenID) error {
	ut := boil.UserToken{
		ID:        int64(id),
		DeletedAt: null.TimeFrom(time.Now()),
	}
	_, err := ut.Update(ctx, r.tx, sb.Infer())
	return err
}
