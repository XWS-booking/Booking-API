package main

import (
	. "auth_service/auth"
	"auth_service/common/messaging"
	"auth_service/common/messaging/nats"
	. "auth_service/database"
	authGrpc "auth_service/proto/auth"
	"fmt"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"auth_service/auth/handlers"
	"auth_service/auth/saga-config"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	QueueGroup = "auth_service"
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
	replySubscriber := initSubscriber(os.Getenv("DELETE_HOST_REPLY_SUBJECT"), QueueGroup)
	commandPublisher := initPublisher(os.Getenv("DELETE_HOST_COMMAND_SUBJECT"))
	replyPublisher := initPublisher(os.Getenv("DELETE_HOST_REPLY_SUBJECT"))
	commandSubscriber := initSubscriber(os.Getenv("DELETE_HOST_COMMAND_SUBJECT"), QueueGroup)

	o := initOrchestator(commandPublisher, replySubscriber)
	authService := &AuthService{UserRepository: userRepository, DeleteHostOrchestrator: o}
	initDeleteHostProfileHandler(authService, replyPublisher, commandSubscriber)

	authController := CreateAuthController(authService)
	authGrpc.RegisterAuthServiceServer(grpcServer, authController)
	fmt.Println("main sub i pub", commandSubscriber, replyPublisher)

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

func initPublisher(subject string) messaging.Publisher {
	publisher, err := nats.NewNATSPublisher(
		os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"),
		os.Getenv("NATS_USER"), os.Getenv("NATS_PASS"), subject)
	if err != nil {
		log.Fatal(err)
	}
	return publisher
}

func initSubscriber(subject, queueGroup string) messaging.Subscriber {
	subscriber, err := nats.NewNATSSubscriber(
		os.Getenv("NATS_HOST"), os.Getenv("NATS_PORT"),
		os.Getenv("NATS_USER"), os.Getenv("NATS_PASS"), subject, queueGroup)
	if err != nil {
		log.Fatal(err)
	}
	return subscriber
}

func initDeleteHostProfileHandler(service *AuthService, publisher messaging.Publisher, subscriber messaging.Subscriber) {
	_, err := handlers.NewDeleteHostCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}

func initOrchestator(publisher messaging.Publisher, subscriber messaging.Subscriber) *saga_config.DeleteHostOrchestrator {
	o, err := saga_config.NewDeleteHostOrchestrator(publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
	return o
}
