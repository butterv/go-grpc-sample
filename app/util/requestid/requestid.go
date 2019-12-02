package requestid

import (
	"context"
	"strings"

	"github.com/rs/xid"
	"google.golang.org/grpc/metadata"
)

const (
	// DefaultXRequestIDKey is a key for getting request id.
	DefaultXRequestIDKey = "X-Request-ID"
	unknownRequestID     = "<unknown>"
)

// GenerateRequestID generates a request id.
func GenerateRequestID() string {
	return xid.New().String()
}

// RequestID takes a request id from metadata.
func RequestID(ctx context.Context) string {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return unknownRequestID
	}

	key := strings.ToLower(DefaultXRequestIDKey)
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
