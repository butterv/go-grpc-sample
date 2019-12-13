package main

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pbv1 "github.com/istsh/go-grpc-sample/app/pb/v1"
	appstatus "github.com/istsh/go-grpc-sample/app/status"
)

func TestDefaultHTTPError(t *testing.T) {
	err := appstatus.Unauthenticated.Err()

	ctx := context.Background()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("", "", nil)
	marshaler := &runtime.JSONPb{}

	HTTPError(ctx, &runtime.ServeMux{}, marshaler, w, req, err)
	if got, want := w.Header().Get("Content-Type"), marshaler.ContentType(); got != want {
		t.Errorf(`w.Header().Get("Content-Type") = %q; want %q`, got, want)
	}
	if got, want := w.Code, http.StatusUnauthorized; got != want {
		t.Errorf("w.Code = %d; want %d", got, want)
	}

	want := pbv1.Error{
		Error: &pbv1.Error_ErrorDetail{
			ErrorCode: "USER_UNAUTHENTICATED",
			Locale:    "ja-JP",
			Message:   "ユーザーの認証ができませんでした。",
		},
	}
	wantBytes, _ := marshaler.Marshal(want)
	if !bytes.Equal(w.Body.Bytes(), wantBytes) {
		t.Errorf("unmatch body: got: %s, want: %s", w.Body.Bytes(), wantBytes)
	}
}

func TestDefaultHTTPError_NoDetails(t *testing.T) {
	errMsg := "not found"
	err := status.New(codes.NotFound, errMsg).Err()

	ctx := context.Background()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("", "", nil)
	marshaler := &runtime.JSONPb{}

	HTTPError(ctx, &runtime.ServeMux{}, marshaler, w, req, err)
	if got, want := w.Header().Get("Content-Type"), marshaler.ContentType(); got != want {
		t.Errorf(`w.Header().Get("Content-Type") = %q; want %q`, got, want)
	}
	if got, want := w.Code, http.StatusNotFound; got != want {
		t.Errorf("w.Code = %d; want %d", got, want)
	}

	want := pbv1.Error{
		Error: &pbv1.Error_ErrorDetail{
			Message: errMsg,
		},
	}
	wantBytes, _ := marshaler.Marshal(want)
	if !bytes.Equal(w.Body.Bytes(), wantBytes) {
		t.Errorf("unmatch body: got: %s, want: %s", w.Body.Bytes(), wantBytes)
	}
}

func TestDefaultHTTPError_ContextType(t *testing.T) {
	ctx := context.Background()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("", "", nil)
	marshaler := &runtime.JSONPb{}

	HTTPError(ctx, &runtime.ServeMux{}, marshaler, w, req, nil)
	if got, want := w.Header().Get("Content-Type"), marshaler.ContentType(); got != want {
		t.Errorf(`w.Header().Get("Content-Type") = %q; want %q`, got, want)
	}
}

func TestDefaultHTTPError_NotStatusError(t *testing.T) {
	err := fmt.Errorf("an error occurred")

	ctx := context.Background()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("", "", nil)
	marshaler := &runtime.JSONPb{}

	HTTPError(ctx, &runtime.ServeMux{}, marshaler, w, req, err)
	if got, want := w.Header().Get("Content-Type"), marshaler.ContentType(); got != want {
		t.Errorf(`w.Header().Get("Content-Type") = %q; want %q`, got, want)
	}
	if got, want := w.Code, http.StatusInternalServerError; got != want {
		t.Errorf("w.Code = %d; want %d", got, want)
	}

	want := pbv1.Error{
		Error: &pbv1.Error_ErrorDetail{
			Message: err.Error(),
		},
	}
	wantBytes, _ := marshaler.Marshal(want)
	if !bytes.Equal(w.Body.Bytes(), wantBytes) {
		t.Errorf("unmatch body: got: %s, want: %s", w.Body.Bytes(), wantBytes)
	}
}
