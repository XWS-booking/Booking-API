package main

import (
	. "accomodation_service/accomodation"
	"accomodation_service/accomodation/services/storage"
	. "accomodation_service/database"
	accomodationGrpc "accomodation_service/proto/accomodation"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	_ "github.com/supabase-community/storage-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
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
	DeclareUnique(db, []UniqueField{})
	if err != nil {
		log.Fatal(err)
		return
	}

	storageService := storage.NewStorageService()
	accomodationRepository := &AccomodationRepository{
		DB: db,
	}
	accomodationService := &AccomodationService{
		AccomodationRepository: accomodationRepository,
	}
	accomodationController := NewAccomodationController(accomodationService, storageService)
	accomodationGrpc.RegisterAccomodationServiceServer(grpcServer, accomodationController)

	// userRepository := &UserRepository{DB: db, Logger: logger}
	// authService := &AuthService{UserRepository: userRepository}
	// authController := CreateAuthController(authService)
	// accomodationGrpc.RegisterAuthServiceServer(grpcServer, authController)

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
