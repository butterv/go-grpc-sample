package interceptor

import (
	"context"

	"google.golang.org/grpc"

	"github.com/istsh/go-grpc-sample/app/util/log"
	"github.com/istsh/go-grpc-sample/app/util/requestid"
)

// RequestIDInterceptor is a interceptor of access control list.
func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := requestid.RequestID(ctx)
		ctx = context.WithValue(ctx, log.CtxRequestIDKey, requestID)
		return handler(ctx, req)
	}
}
