package main

import (
	"context"
	"fmt"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"github.com/sirupsen/logrus"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"os/signal"
	"recommendation_service/database"
	"recommendation_service/messaging"
	"recommendation_service/messaging/nats"
	"recommendation_service/opentelementry"
	recommendationGrpc "recommendation_service/proto/recommendation"
	"recommendation_service/recommendation"
	"recommendation_service/recommendation/repositories/accommodation"
	"recommendation_service/recommendation/repositories/rating"
	"recommendation_service/recommendation/repositories/user"
	"recommendation_service/recommendation/services"
	"syscall"
)

const (
	queueGroup = "RECOMMENDATION_SERVICE"
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
	log.Println("Server started")
	if err != nil {
		fmt.Println("here")
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

	db, err := database.InitDB()

	if err != nil {
		log.Fatal("database connection failed, reason: ", err.Error())
		return
	}

	defer func(db neo4j.DriverWithContext, ctx context.Context) {
		err := db.Close(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(db, context.TODO())

	accommodationRepository := accommodation.AccommodationRepository{Db: db}
	userRepository := user.UserRepository{Db: db}
	ratingRepository := rating.RatingRepository{Db: db}

	accommodationListener := initSubscriber("ACCOMMODATION_EVENT", queueGroup)
	userListener := initSubscriber("USER_EVENT", queueGroup)
	ratingListener := initSubscriber("RATING_EVENT", queueGroup)

	recommendationService := services.RecommendationService{
		AccommodationRepository: &accommodationRepository,
		UserRepository:          &userRepository,
		RatingRepository:        &ratingRepository,
	}

	recommendationController := recommendation.NewRecommendationController(recommendationService, accommodationListener, userListener, ratingListener)
	recommendationGrpc.RegisterRecommendationServiceServer(grpcServer, recommendationController)

	StartServer(grpcServer, listener)
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

func StartServer(grpcServer *grpc.Server, listener net.Listener) {
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
