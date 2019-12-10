package interceptor

import (
	"context"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	"github.com/istsh/go-grpc-sample/app/util/log"
)

const (
	// XRequestIDKey is a key for getting request id.
	XRequestIDKey    = "X-Request-ID"
	unknownRequestID = "<unknown>"
)

// RequestIDInterceptor is a interceptor of access control list.
func RequestIDInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		requestID := requestIDFromContext(ctx)
		ctx = context.WithValue(ctx, log.CtxRequestIDKey, requestID)
		return handler(ctx, req)
	}
}

func requestIDFromContext(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return unknownRequestID
	}

	key := strings.ToLower(XRequestIDKey)
	header, ok := md[key]
	if !ok || len(header) == 0 {
		return unknownRequestID
	}

	requestID := header[0]
	if requestID == "" {
		return unknownRequestID
	}

	return requestID
}
