package main

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/istsh/go-grpc-sample/app/infrastructure/interceptor"
)

func TestRequestIDAnnotator_NewRequestID(t *testing.T) {
	ctx := context.Background()
	req := httptest.NewRequest("GET", "/", nil)

	key := strings.ToLower(interceptor.XRequestIDKey)
	md := RequestIDAnnotator(ctx, req)
	header := md[key]
	if len(header) == 0 {
		t.Fatalf("%s is not found in header", key)
	}
	if header[0] == req.Header.Get(interceptor.XRequestIDKey) {
		t.Error("request ids matched")
	}
}

func TestRequestIDAnnotator_UseHeaderValue(t *testing.T) {
	ctx := context.Background()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(interceptor.XRequestIDKey, "request_id")

	key := strings.ToLower(interceptor.XRequestIDKey)
	md := RequestIDAnnotator(ctx, req)
	header := md[key]
	if len(header) == 0 {
		t.Fatalf("%s is not found in header", key)
	}
	if header[0] != req.Header.Get(interceptor.XRequestIDKey) {
		t.Error("request ids mismatched")
	}
}
