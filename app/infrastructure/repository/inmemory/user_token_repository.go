package inmemory

import (
	"context"
	"database/sql"
	"time"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type inmemoryUserTokenRepository struct {
	repository.UserTokenRepositoryAccess

	s *Store
}

type inmemoryTxUserTokenRepository struct {
	repository.UserTokenRepositoryModify

	s *Store
}

func (r inmemoryUserTokenRepository) FindByToken(ctx context.Context, token string) (*model.UserToken, error) {
	for _, ut := range r.s.userTokens {
		if ut.Token == token {
			return ut, nil
		}
	}

	return nil, nil
}

func (r inmemoryTxUserTokenRepository) Create(ctx context.Context, userID model.UserID, token string) error {
	now := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	ut := &model.UserToken{
		ID:        model.UserTokenID(len(r.s.userTokens) + 1),
		UserID:    userID,
		Token:     token,
		CreatedAt: now,
	}
	r.s.AddUserTokens(ut)

	return nil
}

func (r inmemoryTxUserTokenRepository) Delete(ctx context.Context, id model.UserTokenID) error {
	var userTokens []*model.UserToken

	for _, ut := range r.s.userTokens {
		if ut.ID == id {
			continue
		}
		userTokens = append(userTokens, ut)
	}
	r.s.userTokens = userTokens

	return nil
}
