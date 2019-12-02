package interceptor

import (
	"context"
	"errors"
	"testing"

	"google.golang.org/grpc/metadata"

	"github.com/istsh/go-grpc-sample/app/util/log"
	"github.com/istsh/go-grpc-sample/app/util/requestid"
)

func testRequestID(ctx context.Context, _ interface{}) (interface{}, error) {
	reqID, ok := ctx.Value(log.CtxRequestIDKey).(string)
	if !ok {
		return nil, errors.New("failed to get request id")
	}
	return reqID, nil
}

func TestRequestIDInterceptor(t *testing.T) {
	want := "test_request_id"

	m := map[string]string{
		requestid.DefaultXRequestIDKey: want,
	}
	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(m))
	got, err := RequestIDInterceptor()(ctx, nil, nil, testRequestID)
	if err != nil {
		t.Fatalf("RequestIDInterceptor()()=_, %#v; want nil", err)
	}
	if got.(string) != want {
		t.Errorf("RequestIDInterceptor()()=%s, _; want %s", got.(string), want)
	}
}

func TestRequestIDInterceptor_Unknown(t *testing.T) {
	want := "<unknown>"

	ctx := metadata.NewIncomingContext(context.Background(), metadata.New(nil))
	got, err := RequestIDInterceptor()(ctx, nil, nil, testRequestID)
	if err != nil {
		t.Fatalf("RequestIDInterceptor()()=_, %#v; want nil", err)
	}
	if got.(string) != want {
		t.Errorf("RequestIDInterceptor()()=%s, _; want %s", got.(string), want)
	}
}
