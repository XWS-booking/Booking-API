package main

import (
	. "auth_service/auth"
	"auth_service/auth/handlers"
	"auth_service/auth/saga-config"
	"auth_service/auth/watchers"
	"auth_service/common/messaging"
	"auth_service/common/messaging/nats"
	. "auth_service/database"
	"auth_service/opentelementry"
	authGrpc "auth_service/proto/auth"
	"context"
	"fmt"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	QueueGroup = "auth_service"
)

func main() {
	// OpenTelemetry
	var err error
	opentelementry.Tp, err = opentelementry.InitTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := opentelementry.Tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()
	otel.SetTracerProvider(opentelementry.Tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

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

	userEventPublisher := initPublisher("USER_EVENT")
	watcher := watchers.UserEventWatcher{DB: db, UserRepository: userRepository, Publisher: userEventPublisher}

	o := initOrchestator(commandPublisher, replySubscriber)
	authService := &AuthService{UserRepository: userRepository, DeleteHostOrchestrator: o}
	initDeleteHostProfileHandler(authService, replyPublisher, commandSubscriber)

	authController := CreateAuthController(authService)
	authGrpc.RegisterAuthServiceServer(grpcServer, authController)
	fmt.Println("main sub i pub", commandSubscriber, replyPublisher)

	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	// Start the change stream listener in a goroutine
	wg.Add(1)
	go func() {
		defer wg.Done()
		watcher.StartWatching(ctx)
	}()

	go func() {
		if err := grpcServer.Serve(listener); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	grpcServer.Stop()
	cancel()
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
