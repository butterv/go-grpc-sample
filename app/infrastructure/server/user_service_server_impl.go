package server

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
	pbv1 "github.com/istsh/go-grpc-sample/app/pb/v1"
	appstatus "github.com/istsh/go-grpc-sample/app/status"
	"github.com/istsh/go-grpc-sample/app/usecase"
	"github.com/istsh/go-grpc-sample/app/util/log"
)

const (
	authorizationScheme = "Bearer"
)

type userServiceServer struct {
	r repository.Repository
	u usecase.UserUserCase
}

// NewUserServiceServer creates user service server implementation.
func NewUserServiceServer(r repository.Repository, u usecase.UserUserCase) pbv1.UserServiceServer {
	return &userServiceServer{
		r: r,
		u: u,
	}
}

func (s *userServiceServer) Authenticate(ctx context.Context, req interface{}) (context.Context, error) {
	return ctx, nil
}

func (s *userServiceServer) CreateUser(ctx context.Context, req *pbv1.CreateUserRequest) (*pbv1.CreateUserResponse, error) {
	logger := log.New(ctx)
	con := s.r.NewConnection()

	err := con.RunTransaction(func(tx repository.Transaction) error {
		if err := s.u.CreateUser(ctx, tx, req.GetEmail(), req.GetPassword()); err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		logger.Error(err.Error())
		return nil, appstatus.FailedToCreateUser.Err()
	}

	return &pbv1.CreateUserResponse{}, nil
}
