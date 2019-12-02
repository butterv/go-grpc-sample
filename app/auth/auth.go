package auth

import (
	"context"

	"github.com/istsh/go-grpc-sample/app/entity/model"
)

type authKey int

const (
	// ctxAuthKey is a key to extract Auth* from context.
	ctxAuthKey authKey = iota
)

// Auth is a base class of the `authenticated` status.
type Auth interface {
	UserID() model.UserID
}

// Auth is an authenticated user information.
type auth struct {
	userID model.UserID
}

func NewAuth(userID model.UserID) Auth {
	return &auth{
		userID: userID,
	}
}

func (a *auth) UserID() model.UserID {
	return a.userID
}

// ContextWithAuth set auth to context.
func ContextWithAuth(ctx context.Context, a Auth) context.Context {
	return context.WithValue(ctx, ctxAuthKey, a)
}

// FromContext retrieves Account from context.
func FromContext(ctx context.Context) Auth {
	a, ok := ctx.Value(ctxAuthKey).(Auth)
	if !ok {
		return nil
	}

	return a
}
