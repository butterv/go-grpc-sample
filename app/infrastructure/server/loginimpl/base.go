package loginimpl

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
	loginpb "github.com/istsh/go-grpc-sample/app/pb/v1/login"
	"github.com/istsh/go-grpc-sample/app/usecase"
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
