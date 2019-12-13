package interceptor

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"path"
	"strings"
	"time"

	"github.com/golang/protobuf/jsonpb"
	"github.com/golang/protobuf/proto"
	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/logging"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

var (
	jsonPbMarshaller = &jsonpb.Marshaler{}
)

type jsonpbRequestMarshalleble struct {
	proto.Message
}

type jsonpbResponseMarshalleble struct {
	proto.Message
}

// PayloadUnaryClientInterceptor returns a new unary client interceptor that logs the paylods of requests and responses.
func PayloadUnaryClientInterceptor(entry *logrus.Entry, decider grpc_logging.ClientPayloadLoggingDecider, startTimeFunc func() time.Time, durationFunc func(startTime time.Time) time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if !decider(ctx, method) {
			return invoker(ctx, method, req, reply, cc, opts...)
		}

		startTime := startTimeFunc()
		err := invoker(ctx, method, req, reply, cc, opts...)
		fields, level := newLogFields(ctx, method, req, reply, durationFunc(startTime), err)

		levelLogf(
			entry.WithFields(fields),
			level,
			"client request/response payload logged as grpc.response.content")
		return err
	}
}

func newLogFields(ctx context.Context, method string, reqPbMsg, resPbMsg interface{}, duration time.Duration, err error) (logrus.Fields, logrus.Level) {
	code := grpc_logging.DefaultErrorToCode(err)
	level := grpc_logrus.DefaultCodeToLevel(code)
	durField, durVal := grpc_logrus.DefaultDurationToField(duration)

	fields := logrus.Fields{
		"system":       "grpc",
		"span.kind":    "client",
		"grpc.service": path.Dir(method)[1:],
		"grpc.method":  path.Base(method),
		"grpc.code":    code.String(),
		durField:       durVal,
	}

	if md, ok := metadata.FromOutgoingContext(ctx); ok {
		if values := md.Get("grpcgateway-user-agent"); len(values) > 0 {
			fields["user-agent"] = values[0]
		}
		if values := md.Get("x-forwarded-for"); len(values) > 0 {
			fields["remote_ip"] = values[0]
		}
		if values := md.Get(XRequestIDKey); len(values) > 0 {
			fields["request_id"] = values[0]
		}
	}
	if p, ok := reqPbMsg.(proto.Message); ok {
		fields["grpc.request.content"] = &jsonpbRequestMarshalleble{p}
	}
	if p, ok := resPbMsg.(proto.Message); ok {
		fields["grpc.response.content"] = &jsonpbResponseMarshalleble{p}
	}
	if err != nil {
		fields[logrus.ErrorKey] = err.Error()
	}

	return fields, level
}

func levelLogf(entry *logrus.Entry, level logrus.Level, format string, args ...interface{}) {
	switch level {
	case logrus.DebugLevel:
		entry.Debugf(format, args...)
	case logrus.InfoLevel:
		entry.Infof(format, args...)
	case logrus.WarnLevel:
		entry.Warningf(format, args...)
	case logrus.ErrorLevel:
		entry.Errorf(format, args...)
	case logrus.FatalLevel:
		entry.Fatalf(format, args...)
	case logrus.PanicLevel:
		entry.Panicf(format, args...)
	}
}

func (j *jsonpbRequestMarshalleble) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	if err := jsonPbMarshaller.Marshal(b, j.Message); err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}
	return filterParam(b.Bytes()), nil
}

func (j *jsonpbResponseMarshalleble) MarshalJSON() ([]byte, error) {
	b := &bytes.Buffer{}
	if err := jsonPbMarshaller.Marshal(b, j.Message); err != nil {
		return nil, fmt.Errorf("jsonpb serializer failed: %v", err)
	}
	return b.Bytes(), nil
}

func filterParam(bs []byte) []byte {
	reqJSON := make(map[string]interface{})
	json.Unmarshal(bs, &reqJSON) // ignore error here.

	var toFilter []string
	for k := range reqJSON {
		if strings.Contains(strings.ToLower(k), "password") {
			toFilter = append(toFilter, k)
		}
	}
	for _, k := range toFilter {
		reqJSON[k] = "<FILTERED>"
	}

	r, _ := json.Marshal(reqJSON) // ignore error here
	return r
}
