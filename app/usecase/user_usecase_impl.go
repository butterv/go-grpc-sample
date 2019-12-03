package usecase

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/rs/xid"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"

	"github.com/istsh/go-grpc-sample/app/entity/model"
	"github.com/istsh/go-grpc-sample/app/entity/repository"
)

type userUserCase struct {
}

func NewUserUsecase() UserUserCase {
	return &userUserCase{}
}

func (uc *userUserCase) CreateUser(ctx context.Context, tx repository.Transaction, email, password string) error {
	userID := model.UserID(xid.New().String())

	if err := tx.User().Create(userID, email); err != nil {
		return err
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return tx.UserPassword().Create(userID, string(hashedPassword))
}

func (uc *userUserCase) IsCorrectUserPassword(ctx context.Context, con repository.Connection, userID model.UserID, password string) (bool, error) {
	up, err := con.UserPassword().Find(userID)
	if err != nil {
		return false, err
	}
	if up == nil {
		return false, fmt.Errorf("password not found (userid: %s)", userID)
	}
	if err := bcrypt.CompareHashAndPassword([]byte(up.PasswordHash), []byte(password)); err != nil {
		return false, nil
	}

	return true, nil
}

func (uc *userUserCase) CreateUserToken(ctx context.Context, tx repository.Transaction, userID model.UserID) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	now := time.Now()
	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = string(userID)
	claims["iat"] = now.Unix()
	claims["exp"] = now.Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
	if err != nil {
		return "", err
	}

	err = tx.UserToken().Create(userID, tokenString)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
