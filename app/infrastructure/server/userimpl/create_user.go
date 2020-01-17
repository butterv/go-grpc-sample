package userimpl

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
	userpb "github.com/istsh/go-grpc-sample/app/pb/v1/user"
	appstatus "github.com/istsh/go-grpc-sample/app/status"
	"github.com/istsh/go-grpc-sample/app/util/log"
)

func (s *userServiceServer) CreateUser(ctx context.Context, req *userpb.CreateUserRequest) (*userpb.CreateUserResponse, error) {
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

	return &userpb.CreateUserResponse{}, nil
}
