package main

import (
	"context"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/istsh/go-grpc-sample/app/util/requestid"
)

func TestRequestIDAnnotator_NewRequestID(t *testing.T) {
	ctx := context.Background()
	req := httptest.NewRequest("GET", "/", nil)

	key := strings.ToLower(requestid.DefaultXRequestIDKey)
	md := RequestIDAnnotator(ctx, req)
	header := md[key]
	if len(header) == 0 {
		t.Fatalf("%s is not found in header", key)
	}
	if header[0] == req.Header.Get(requestid.DefaultXRequestIDKey) {
		t.Error("request ids matched")
	}
}

func TestRequestIDAnnotator_UseHeaderValue(t *testing.T) {
	ctx := context.Background()
	req := httptest.NewRequest("GET", "/", nil)
	req.Header.Set(requestid.DefaultXRequestIDKey, "request_id")

	key := strings.ToLower(requestid.DefaultXRequestIDKey)
	md := RequestIDAnnotator(ctx, req)
	header := md[key]
	if len(header) == 0 {
		t.Fatalf("%s is not found in header", key)
	}
	if header[0] != req.Header.Get(requestid.DefaultXRequestIDKey) {
		t.Error("request ids mismatched")
	}
}
