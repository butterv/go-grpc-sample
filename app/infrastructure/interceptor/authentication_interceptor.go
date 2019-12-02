package interceptor

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// Authenticator provides Authenticate method.
// Each service should implement this interface, otherwise, all requests will be rejected with authentication error.
type Authenticator interface {
	Authenticate(ctx context.Context, req interface{}) (context.Context, error)
}

// AuthenticationInterceptor is a interceptor of authentication.
func AuthenticationInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		authenticator, ok := info.Server.(Authenticator)
		if !ok {
			// If Service doesn't implement Authenticator, return InternalServerError always.
			return nil, status.New(codes.Internal, "Authenticator is not implemented").Err()
		}

		ctx, err := authenticator.Authenticate(ctx, req)
		if err != nil {
			return nil, status.New(codes.Unauthenticated, fmt.Sprintf("Not authenticated: %v", err)).Err()
		}

		return handler(ctx, req)
	}
}
