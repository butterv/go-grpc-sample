package main

import (
	"context"
	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/golang/glog"
	"github.com/gorilla/handlers"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/encoding/gzip"

	"github.com/istsh/go-grpc-sample/app/infrastructure/interceptor"
	loginpb "github.com/istsh/go-grpc-sample/app/pb/v1/login"
	userpb "github.com/istsh/go-grpc-sample/app/pb/v1/user"
)

var (
	grpcServerEndpoint = flag.String("grpc-server-endpoint", "localhost:9090", "gRPC server endpoint")
)

func grpcDialOptions() []grpc.DialOption {
	l := logrus.New()
	l.SetFormatter(&logrus.JSONFormatter{})
	decider := func(ctx context.Context, fullMethodName string) bool {
		return true
	}
	startTimeFunc := func() time.Time {
		return time.Now()
	}
	durationFunc := func(startTime time.Time) time.Duration {
		return time.Now().Sub(startTime)
	}

	opts := []grpc.DialOption{
		grpc.WithInsecure(),
		grpc.WithDefaultCallOptions(grpc.UseCompressor(gzip.Name)),
		// Output request/response payload logs
		grpc.WithUnaryInterceptor(interceptor.PayloadUnaryClientInterceptor(logrus.NewEntry(l), decider, startTimeFunc, durationFunc)),
	}

	return opts
}

func registerServiceHandlers(ctx context.Context, mux *runtime.ServeMux) error {
	opts := grpcDialOptions()

	if err := loginpb.RegisterLoginServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	if err := userpb.RegisterUserServiceHandlerFromEndpoint(ctx, mux, *grpcServerEndpoint, opts); err != nil {
		return err
	}

	return nil
}

func run() error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	runtime.HTTPError = HTTPError

	mux := runtime.NewServeMux(
		// Set request_id to grpc metadata
		runtime.WithMetadata(RequestIDAnnotator),
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{OrigName: false, EmitDefaults: true}),
	)

	if err := registerServiceHandlers(ctx, mux); err != nil {
		return err
	}

	handler := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{http.MethodPost, http.MethodGet, http.MethodPut, http.MethodDelete}),
		handlers.AllowedHeaders([]string{"Authorization", "Content-Type", "Accept-Encoding", "Accept"}),
	)(mux)

	addr := ":8080"
	fmt.Printf("http server started on %s\n", addr)
	// Start HTTP server (and proxy calls to gRPC server endpoint)
	return http.ListenAndServe(addr, handler)
}

func main() {
	flag.Parse()
	defer glog.Flush()

	if err := run(); err != nil {
		glog.Fatal(err)
	}
}
