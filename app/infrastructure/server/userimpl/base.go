package userimpl

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
	userpb "github.com/istsh/go-grpc-sample/app/pb/v1/user"
	"github.com/istsh/go-grpc-sample/app/usecase"
)

type userServiceServer struct {
	r repository.Repository
	u usecase.UserUserCase
}

// NewUserServiceServer creates user service server implementation.
func NewUserServiceServer(r repository.Repository, u usecase.UserUserCase) userpb.UserServiceServer {
	return &userServiceServer{
		r: r,
		u: u,
	}
}

func (s *userServiceServer) Authenticate(ctx context.Context, req interface{}) (context.Context, error) {
	return ctx, nil
}
