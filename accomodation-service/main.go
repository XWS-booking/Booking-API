package main

import (
	. "accomodation_service/accomodation"
	"accomodation_service/accomodation/handlers"
	"accomodation_service/accomodation/services/storage"
	"accomodation_service/accomodation/watchers"
	"accomodation_service/common/messaging"
	"accomodation_service/common/messaging/nats"
	. "accomodation_service/database"
	"accomodation_service/opentelementry"
	accomodationGrpc "accomodation_service/proto/accomodation"
	"context"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	_ "github.com/supabase-community/storage-go"
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
	QueueGroup = "accommodation_service"
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
	commandSubscriber := initSubscriber(os.Getenv("DELETE_HOST_COMMAND_SUBJECT"), QueueGroup)
	replyPublisher := initPublisher(os.Getenv("DELETE_HOST_REPLY_SUBJECT"))

	accommodationEventPublisher := initPublisher("ACCOMMODATION_EVENT")
	watcher := watchers.AccommodationEventWatcher{
		DB:                      db,
		AccommodationRepository: accomodationRepository,
		Publisher:               accommodationEventPublisher,
	}
	initDeleteHostProfileHandler(accomodationService, replyPublisher, commandSubscriber)

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
	cancel()
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

func initDeleteHostProfileHandler(service *AccomodationService, publisher messaging.Publisher, subscriber messaging.Subscriber) {
	_, err := handlers.NewDeleteHostCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
