package main

import (
	"net"

	_ "github.com/go-sql-driver/mysql"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/jinzhu/gorm"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	_ "google.golang.org/grpc/encoding/gzip"
	"google.golang.org/grpc/reflection"

	"github.com/istsh/go-grpc-sample/app/entity/repository"
	"github.com/istsh/go-grpc-sample/app/infrastructure/interceptor"
	"github.com/istsh/go-grpc-sample/app/infrastructure/repository/persistence"
	"github.com/istsh/go-grpc-sample/app/infrastructure/server"
	loginpb "github.com/istsh/go-grpc-sample/app/pb/v1/login"
	userpb "github.com/istsh/go-grpc-sample/app/pb/v1/user"
	"github.com/istsh/go-grpc-sample/app/usecase"
)

func connectDB() *gorm.DB {
	db, err := gorm.Open("mysql", "root:password@tcp(db)/sample?charset=utf8mb4&parseTime=True&loc=UTC")
	if err != nil {
		panic(err.Error())
	}
	db.LogMode(true)
	return db
}

func newGRPCServer(r repository.Repository, u usecase.UserUserCase) *grpc.Server {
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			interceptor.RequestIDInterceptor(),
			interceptor.AuthenticationInterceptor(),
			grpc_validator.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
	)

	loginpb.RegisterLoginServiceServer(s, server.NewLoginServiceServer(r, u))
	userpb.RegisterUserServiceServer(s, server.NewUserServiceServer(r, u))

	return s
}

func main() {
	db := connectDB()
	defer db.Close()

	r := persistence.NewDBRepository(db)
	u := usecase.NewUserUsecase()

	listenPort, err := net.Listen("tcp", ":9090")
	if err != nil {
		logrus.Fatalln(err)
	}

	s := newGRPCServer(r, u)
	reflection.Register(s)
	s.Serve(listenPort)
}
