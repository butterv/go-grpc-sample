package interceptor

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"testing"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type testAuthenticationServer struct {
	ctx context.Context
	err error
}

func (t *testAuthenticationServer) Authenticate(context.Context, interface{}) (context.Context, error) {
	return t.ctx, t.err
}

func testAuthenticationHandler(context.Context, interface{}) (interface{}, error) {
	return nil, nil
}

func TestAuthenticationInterceptor(t *testing.T) {
	server := &testAuthenticationServer{
		ctx: context.Background(),
		err: nil,
	}
	info := &grpc.UnaryServerInfo{
		Server: server,
	}

	_, err := AuthenticationInterceptor()(context.Background(), nil, info, testAuthenticationHandler)
	if err != nil {
		t.Fatalf("AuthenticationInterceptor()()=_, %#v; want nil", err)
	}
}

func TestAuthenticationInterceptor_AuthenticatorNotImplemented(t *testing.T) {
	wantErr := status.New(codes.Internal, "Authenticator is not implemented").Err()

	info := &grpc.UnaryServerInfo{
		Server: struct{}{},
	}

	_, err := AuthenticationInterceptor()(context.Background(), nil, info, testAuthenticationHandler)
	if err == nil {
		t.Fatalf("AuthenticationInterceptor()()=_, nil; want %#v", wantErr)
	}
	if !reflect.DeepEqual(err, wantErr) {
		t.Errorf("AuthenticationInterceptor()()=_, %#v; want %#v", err, wantErr)
	}
}

func TestAuthenticationInterceptor_Unauthenticated(t *testing.T) {
	wantErr := status.New(codes.Unauthenticated, fmt.Sprintf("Not authenticated: %v", errors.New("an error occurred"))).Err()

	server := &testAuthenticationServer{
		ctx: context.Background(),
		err: errors.New("an error occurred"),
	}
	info := &grpc.UnaryServerInfo{
		Server: server,
	}

	_, err := AuthenticationInterceptor()(context.Background(), nil, info, testAuthenticationHandler)
	if err == nil {
		t.Fatalf("AuthenticationInterceptor()()=_, nil; want %#v", wantErr)
	}
	if !reflect.DeepEqual(err, wantErr) {
		t.Errorf("AuthenticationInterceptor()()=_, %#v; want %#v", err, wantErr)
	}
}
