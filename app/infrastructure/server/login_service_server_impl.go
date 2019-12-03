package server

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
	loginpb "github.com/istsh/go-grpc-sample/app/pb/v1/login"
	appstatus "github.com/istsh/go-grpc-sample/app/status"
	"github.com/istsh/go-grpc-sample/app/usecase"
	"github.com/istsh/go-grpc-sample/app/util/log"
)

type loginServiceServer struct {
	r repository.Repository
	u usecase.UserUserCase
}

// NewLoginServiceServer creates login service server implementation.
func NewLoginServiceServer(r repository.Repository, u usecase.UserUserCase) loginpb.LoginServiceServer {
	return &loginServiceServer{
		r: r,
		u: u,
	}
}

func (s *loginServiceServer) Authenticate(ctx context.Context, req interface{}) (context.Context, error) {
	return ctx, nil
}

func (s *loginServiceServer) Login(ctx context.Context, req *loginpb.LoginRequest) (*loginpb.LoginResponse, error) {
	logger := log.New(ctx)
	con := s.r.NewConnection()

	u, err := con.User().FindByEmail(req.GetEmail())
	if err != nil {
		logger.Error(err.Error())
		return nil, appstatus.FailedToLogin.Err()
	}
	if u == nil {
		return nil, appstatus.FailedToLogin.Err()
	}

	isCorrect, err := s.u.IsCorrectUserPassword(ctx, con, u.ID, req.GetPassword())
	if err != nil {
		logger.Error(err.Error())
		return nil, appstatus.FailedToLogin.Err()
	}
	if !isCorrect {
		return nil, appstatus.FailedToLogin.Err()
	}

	var tokenString string
	err = con.RunTransaction(func(tx repository.Transaction) error {
		token, err := s.u.CreateUserToken(ctx, tx, u.ID)
		if err != nil {
			return err
		}

		tokenString = token
		return nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, appstatus.FailedToLogin.Err()
	}

	return &loginpb.LoginResponse{
		Token: tokenString,
	}, nil
}
