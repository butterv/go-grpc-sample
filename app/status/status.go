package status

import (
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	errorpb "github.com/istsh/go-grpc-sample/app/pb/v1/error"
)

const (
	// LocaleJaJp is a Japanese locale id in Japan.
	LocaleJaJp = "ja-JP"
	// LocaleEnUs is a English locale id in United States.
	LocaleEnUs = "en-US"
)

func must(s *status.Status, err error) *status.Status {
	if err != nil {
		panic(err)
	}
	return s
}

var (
	// AnErrorHasOccurred is a gRPC status when an error occurs.
	AnErrorHasOccurred = must(status.New(codes.Internal, "an error occurred").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "AN_ERROR_HAS_OCCURRED",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJaJp,
			Message: "エラーが発生しました。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUs,
			Message: "an error has occurred",
		},
	))

	// Unauthenticated is a gRPC status when `auth` cannot be gotten from the context.
	Unauthenticated = must(status.New(codes.Unauthenticated, "not authenticated").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "USER_UNAUTHENTICATED",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJaJp,
			Message: "ユーザーの認証ができませんでした。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUs,
			Message: "Unauthenticated",
		},
	))

	//
	FailedToCreateUser = must(status.New(codes.Internal, "failed to create user").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "FAILED_TO_CREATE_USER",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJaJp,
			Message: "ユーザー登録に失敗しました。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUs,
			Message: "Failed to create user",
		},
	))

	//
	FailedToChangePassword = must(status.New(codes.Internal, "failed to change password").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "FAILED_TO_CHANGE_PASSWORD",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJaJp,
			Message: "パスワードの変更に失敗しました。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUs,
			Message: "Failed to change password",
		},
	))

	// NoSuchUser is a gRPC status that indicates the specified user isn't found.
	NoSuchUser = must(status.New(codes.InvalidArgument, "no such a user").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "NO_SUCH_USER",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJaJp,
			Message: "ユーザーが見つかりませんでした。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUs,
			Message: "No such user",
		},
	))

	// FailedToLogin is a gRPC status when user failed to login.
	FailedToLogin = must(status.New(codes.Unauthenticated, "failed to login").WithDetails(
		&errorpb.ErrorCode{
			ErrorCode: "USER_FAILED_TO_LOGIN",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleJaJp,
			Message: "ログインできませんでした。",
		},
		&errdetails.LocalizedMessage{
			Locale:  LocaleEnUs,
			Message: "Failed to login",
		},
	))
)

// PickStatusError returns err if err is statusError, otherwise it returns the second argument error.
func PickStatusError(err error, st *status.Status) error {
	if st == nil {
		return AnErrorHasOccurred.Err()
	}
	if _, ok := status.FromError(err); ok {
		return err
	}
	return st.Err()
}
