package main

import (
	"context"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"

	"github.com/istsh/go-grpc-sample/app/util/requestid"
)

// RequestIDAnnotator takes requestID from http request header and sets it to metadata.
func RequestIDAnnotator(ctx context.Context, req *http.Request) metadata.MD {
	requestID := req.Header.Get(requestid.DefaultXRequestIDKey)
	if requestID == "" {
		requestID = requestid.GenerateRequestID()
	}

	key := strings.ToLower(requestid.DefaultXRequestIDKey)
	return metadata.New(map[string]string{
		key: requestID,
	})
}
