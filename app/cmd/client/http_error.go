package main

import (
	"context"
	"io"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	pbv1 "github.com/istsh/go-grpc-sample/app/pb/v1"
	appstatus "github.com/istsh/go-grpc-sample/app/status"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/grpclog"
	"google.golang.org/grpc/status"
)

// HTTPError is the default implementation of HTTPError.
// See. https://github.com/grpc-ecosystem/grpc-gateway/blob/master/runtime/errors.go#L89
// Except for a part, it is the same code as runtime.DefaultHTTPError.
func HTTPError(_ context.Context, _ *runtime.ServeMux, marshaler runtime.Marshaler, w http.ResponseWriter, _ *http.Request, err error) {
	const fallback = `{"error": "failed to marshal error message"}`

	s, ok := status.FromError(err)
	if !ok {
		s = status.New(codes.Unknown, err.Error())
	}

	w.Header().Del("Trailer")

	contentType := marshaler.ContentType()
	if httpBodyMarshaler, ok := marshaler.(*runtime.HTTPBodyMarshaler); ok {
		pb := s.Proto()
		contentType = httpBodyMarshaler.ContentTypeFromMessage(pb)
	}
	w.Header().Set("Content-Type", contentType)

	// This block is the original.
	ed := &pbv1.Error_ErrorDetail{}
	if len(s.Details()) > 0 {
		for _, detail := range s.Details() {
			switch v := detail.(type) {
			case *pbv1.ErrorCode:
				ed.ErrorCode = v.GetErrorCode()
			case *errdetails.LocalizedMessage:
				if ed.GetMessage() != "" {
					// Already set error message.
					continue
				}
				if v.GetLocale() == appstatus.LocaleJaJp {
					ed.Locale = v.GetLocale()
					ed.Message = v.GetMessage()
				}
			}
		}
	} else {
		ed.Message = s.Message()
	}
	e := pbv1.Error{
		Error: ed,
	}

	buf, merr := marshaler.Marshal(e)
	if merr != nil {
		grpclog.Infof("Failed to marshal error message %q: %v", e, merr)
		w.WriteHeader(http.StatusInternalServerError)
		if _, err := io.WriteString(w, fallback); err != nil {
			grpclog.Infof("Failed to write response: %v", err)
		}
		return
	}

	st := runtime.HTTPStatusFromCode(s.Code())
	w.WriteHeader(st)
	if _, err := w.Write(buf); err != nil {
		grpclog.Infof("Failed to write response: %v", err)
	}
}
