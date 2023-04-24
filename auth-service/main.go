package main

import (
	. "auth_service/auth"
	. "auth_service/database"
	authGrpc "auth_service/proto/auth"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	listener, err := net.Listen("tcp", ":"+os.Getenv("PORT"))

	if err != nil {
		log.Fatalln(err)
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(listener)

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(CreateServerLogger()),
	)
	reflection.Register(grpcServer)

	db, err := InitDB()
	DeclareUnique(db, []UniqueField{
		{Collection: "users", Fields: []string{"email"}},
	})
	if err != nil {
		log.Fatal(err)
		return
	}
	logger := log.New(os.Stdout, "[Users-api] ", log.LstdFlags)
	userRepository := &UserRepository{DB: db, Logger: logger}
	authService := &AuthService{UserRepository: userRepository}
	authController := CreateAuthController(authService)
	authGrpc.RegisterAuthServiceServer(grpcServer, authController)

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
}
func CreateServerLogger() grpc.UnaryServerInterceptor {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	entry := logrus.NewEntry(logger)
	return grpc_logrus.UnaryServerInterceptor(entry, grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel))
}
