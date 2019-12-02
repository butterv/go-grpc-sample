package interceptor

import (
	"context"
	"errors"
	"path"
	"reflect"
	"testing"
	"time"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"

	appstatus "github.com/istsh/go-grpc-sample/app/status"
)

type args struct {
	ctx       context.Context
	decided   bool
	startTime time.Time
	duration  time.Duration
	err       error
}

func (a *args) decider(context.Context, string) bool {
	return a.decided
}

func (a *args) invoker(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, opts ...grpc.CallOption) error {
	return a.err
}

func (a *args) getStartTime() time.Time {
	return a.startTime
}

func (a *args) getDuration(time.Time) time.Duration {
	return a.duration
}

func testLogrusEntry() *logrus.Entry {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	return logrus.NewEntry(l)
}

func TestPayloadUnaryClientInterceptor_DeciderReturnFalse(t *testing.T) {
	a := args{
		decided: false,
		err:     nil,
	}
	err := PayloadUnaryClientInterceptor(nil, a.decider, a.getStartTime, a.getDuration)(a.ctx, "", nil, nil, nil, a.invoker)
	if err != nil {
		t.Fatalf("PayloadUnaryClientInterceptor()()=%#v; want nil", err)
	}
}

func TestPayloadUnaryClientInterceptor_DeciderReturnFalse_Error(t *testing.T) {
	wantErr := errors.New("an error occurred")

	a := args{
		decided: false,
		err:     wantErr,
	}
	err := PayloadUnaryClientInterceptor(nil, a.decider, a.getStartTime, a.getDuration)(a.ctx, "", nil, nil, nil, a.invoker)
	if err == nil {
		t.Fatalf("PayloadUnaryClientInterceptor()()=nil; want %#v", err)
	}
	if !reflect.DeepEqual(err, wantErr) {
		t.Errorf("PayloadUnaryClientInterceptor()()=%#v; want %#v", err, wantErr)
	}
}

func TestNewLogFields_Simple(t *testing.T) {
	ctx := context.Background()
	method := "/adminservice.AdminService/CreateAdmin"
	duration := time.Duration(1000)
	durField, durVal := grpc_logrus.DefaultDurationToField(duration)
	code := grpc_logging.DefaultErrorToCode(nil)
	wantLevel := grpc_logrus.DefaultCodeToLevel(code)

	wantFields := logrus.Fields{
		"system":       "grpc",
		"span.kind":    "client",
		"grpc.service": path.Dir(method)[1:],
		"grpc.method":  path.Base(method),
		"grpc.code":    code.String(),
		durField:       durVal,
	}

	gotFields, gotLevel := newLogFields(ctx, method, nil, nil, duration, nil)
	if !reflect.DeepEqual(gotFields, wantFields) {
		t.Errorf("newLogFields()=%#v, _; want %#v", gotFields, wantFields)
	}
	if !reflect.DeepEqual(gotLevel, wantLevel) {
		t.Errorf("newLogFields()=_, %#v; want %#v", gotLevel, wantLevel)
	}
}

func TestNewLogFields_IncludeError(t *testing.T) {
	err := appstatus.Unauthenticated.Err()

	ctx := context.Background()
	method := "/adminservice.AdminService/CreateAdmin"
	duration := time.Duration(1000)
	durField, durVal := grpc_logrus.DefaultDurationToField(duration)
	code := grpc_logging.DefaultErrorToCode(err)
	wantLevel := grpc_logrus.DefaultCodeToLevel(code)

	wantFields := logrus.Fields{
		"system":        "grpc",
		"span.kind":     "client",
		"grpc.service":  path.Dir(method)[1:],
		"grpc.method":   path.Base(method),
		"grpc.code":     code.String(),
		durField:        durVal,
		logrus.ErrorKey: err.Error(),
	}

	gotFields, gotLevel := newLogFields(ctx, method, nil, nil, duration, err)
	if !reflect.DeepEqual(gotFields, wantFields) {
		t.Errorf("newLogFields()=%#v, _; want %#v", gotFields, wantFields)
	}
	if !reflect.DeepEqual(gotLevel, wantLevel) {
		t.Errorf("newLogFields()=_, %#v; want %#v", gotLevel, wantLevel)
	}
}

func TestNewLogFields_IncludeUserAgent(t *testing.T) {
	userAgent := "TEST_USER_AGENT"

	m := map[string]string{
		"grpcgateway-user-agent": userAgent,
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(m))
	method := "/adminservice.AdminService/CreateAdmin"
	duration := time.Duration(1000)
	durField, durVal := grpc_logrus.DefaultDurationToField(duration)
	code := grpc_logging.DefaultErrorToCode(nil)
	wantLevel := grpc_logrus.DefaultCodeToLevel(code)

	wantFields := logrus.Fields{
		"system":       "grpc",
		"span.kind":    "client",
		"grpc.service": path.Dir(method)[1:],
		"grpc.method":  path.Base(method),
		"grpc.code":    code.String(),
		durField:       durVal,
		"user-agent":   userAgent,
	}

	gotFields, gotLevel := newLogFields(ctx, method, nil, nil, duration, nil)
	if !reflect.DeepEqual(gotFields, wantFields) {
		t.Errorf("newLogFields()=%#v, _; want %#v", gotFields, wantFields)
	}
	if !reflect.DeepEqual(gotLevel, wantLevel) {
		t.Errorf("newLogFields()=_, %#v; want %#v", gotLevel, wantLevel)
	}
}

func TestNewLogFields_IncludeXForwardedFor(t *testing.T) {
	xForwardedFor := "TEST_X_FORWARDED_FOR"

	m := map[string]string{
		"x-forwarded-for": xForwardedFor,
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(m))
	method := "/adminservice.AdminService/CreateAdmin"
	duration := time.Duration(1000)
	durField, durVal := grpc_logrus.DefaultDurationToField(duration)
	code := grpc_logging.DefaultErrorToCode(nil)
	wantLevel := grpc_logrus.DefaultCodeToLevel(code)

	wantFields := logrus.Fields{
		"system":       "grpc",
		"span.kind":    "client",
		"grpc.service": path.Dir(method)[1:],
		"grpc.method":  path.Base(method),
		"grpc.code":    code.String(),
		durField:       durVal,
		"remote_ip":    xForwardedFor,
	}

	gotFields, gotLevel := newLogFields(ctx, method, nil, nil, duration, nil)
	if !reflect.DeepEqual(gotFields, wantFields) {
		t.Errorf("newLogFields()=%#v, _; want %#v", gotFields, wantFields)
	}
	if !reflect.DeepEqual(gotLevel, wantLevel) {
		t.Errorf("newLogFields()=_, %#v; want %#v", gotLevel, wantLevel)
	}
}

func TestNewLogFields_IncludeXRequestID(t *testing.T) {
	xRequestID := "TEST_X_REQUEST_ID"

	m := map[string]string{
		"x-request-id": xRequestID,
	}
	ctx := metadata.NewOutgoingContext(context.Background(), metadata.New(m))
	method := "/adminservice.AdminService/CreateAdmin"
	duration := time.Duration(1000)
	durField, durVal := grpc_logrus.DefaultDurationToField(duration)
	code := grpc_logging.DefaultErrorToCode(nil)
	wantLevel := grpc_logrus.DefaultCodeToLevel(code)

	wantFields := logrus.Fields{
		"system":       "grpc",
		"span.kind":    "client",
		"grpc.service": path.Dir(method)[1:],
		"grpc.method":  path.Base(method),
		"grpc.code":    code.String(),
		durField:       durVal,
		"request_id":   xRequestID,
	}

	gotFields, gotLevel := newLogFields(ctx, method, nil, nil, duration, nil)
	if !reflect.DeepEqual(gotFields, wantFields) {
		t.Errorf("newLogFields()=%#v, _; want %#v", gotFields, wantFields)
	}
	if !reflect.DeepEqual(gotLevel, wantLevel) {
		t.Errorf("newLogFields()=_, %#v; want %#v", gotLevel, wantLevel)
	}
}
