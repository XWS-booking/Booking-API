package main

import (
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	_ "github.com/supabase-community/storage-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"reservation_service/common/messaging"
	"reservation_service/common/messaging/nats"
	. "reservation_service/database"
	reservationGrpc "reservation_service/proto/reservation"
	. "reservation_service/reservation"
	"reservation_service/reservation/handlers"
	"syscall"
)

const (
	QueueGroup = "reservation_service"
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

	reservationRepository := &ReservationRepository{
		DB: db,
	}
	reservationService := &ReservationService{
		ReservationRepository: reservationRepository,
	}
	reservationController := NewReservationController(reservationService)
	reservationGrpc.RegisterReservationServiceServer(grpcServer, reservationController)

	commandSubscriber := initSubscriber(os.Getenv("DELETE_HOST_COMMAND_SUBJECT"), QueueGroup)
	replyPublisher := initPublisher(os.Getenv("DELETE_HOST_REPLY_SUBJECT"))
	initDeleteHostProfileHandler(reservationService, replyPublisher, commandSubscriber)

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

func initDeleteHostProfileHandler(service *ReservationService, publisher messaging.Publisher, subscriber messaging.Subscriber) {
	_, err := handlers.NewDeleteHostCommandHandler(service, publisher, subscriber)
	if err != nil {
		log.Fatal(err)
	}
}
