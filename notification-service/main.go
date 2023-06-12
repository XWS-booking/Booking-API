package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	_ "github.com/supabase-community/storage-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	. "notification_service/database"
	. "notification_service/notification"
	"notification_service/opentelementry"
	notificationGrpc "notification_service/proto/notification"
	"os"
	"os/signal"
	"syscall"
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

	startRedisConnection()
	go runGRPCServer()
	fmt.Println("Server started")

	stopCh := make(chan os.Signal, 1)
	signal.Notify(stopCh, syscall.SIGINT, syscall.SIGTERM)
	<-stopCh

	fmt.Println("Server stopped")
}

func runGRPCServer() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal("gRPC server error: ", err)
	}

	grpcServer := grpc.NewServer(
		grpc.UnaryInterceptor(CreateServerLogger()),
	)
	reflection.Register(grpcServer)

	db, err := InitDB()
	DeclareUnique(db, []UniqueField{})
	if err != nil {
		log.Fatal("Database error: ", err)
	}

	notificationRepository := &NotificationRepository{
		DB: db,
	}
	notificationService := &NotificationService{
		NotificationRepository: notificationRepository,
	}

	notificationController := NewNotificationController(notificationService)
	notificationGrpc.RegisterNotificationServiceServer(grpcServer, notificationController)

	err = grpcServer.Serve(listener)
	if err != nil {
		log.Fatal("gRPC server error: ", err)
	}
}

func startRedisConnection() {
	options, err := redis.ParseURL("redis://default:f8KuO1PmiIpo4XS5tKUxZ7DbW2YO4JoP@redis-16014.c300.eu-central-1-1.ec2.cloud.redislabs.com:16014")
	if err != nil {
		log.Fatal("Failed to parse REDIS_URL:", err)
	}
	//Client = redis.NewClient(&redis.Options{
	//	Addr: "host.docker.internal:6379",
	//})
	Client = redis.NewClient(options)
	pong, err := Client.Ping(context.Background()).Result()
	if err != nil {
		log.Fatal("Failed to connect to Redis:", err)
	}
	fmt.Println("Connected to Redis:", pong)
}

func CreateServerLogger() grpc.UnaryServerInterceptor {
	logger := logrus.New()
	logger.SetLevel(logrus.InfoLevel)
	entry := logrus.NewEntry(logger)
	return grpc_logrus.UnaryServerInterceptor(entry, grpc_logrus.WithLevels(grpc_logrus.DefaultCodeToLevel))
}
