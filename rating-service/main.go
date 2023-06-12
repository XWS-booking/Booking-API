package main

import (
	"context"
	"github.com/gin-gonic/gin"
	grpc_logrus "github.com/grpc-ecosystem/go-grpc-middleware/logging/logrus"
	"github.com/sirupsen/logrus"
	_ "github.com/supabase-community/storage-go"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	. "rating_service/database"
	"rating_service/opentelementry"
	ratingGrpc "rating_service/proto/rating"
	. "rating_service/rating"
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

	ratingRepository := &RatingRepository{
		DB: db,
	}
	ratingService := &RatingService{
		RatingRepository: ratingRepository,
	}
	ratingController := NewRatingController(ratingService)
	ratingGrpc.RegisterRatingServiceServer(grpcServer, ratingController)

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

func httpErrorBadRequest(err error, span trace.Span, ctx *gin.Context) {
	httpError(err, span, ctx, http.StatusBadRequest)
}

func httpErrorInternalServerError(err error, span trace.Span, ctx *gin.Context) {
	httpError(err, span, ctx, http.StatusInternalServerError)
}

func httpError(err error, span trace.Span, ctx *gin.Context, status int) {
	log.Println(err.Error())
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
	ctx.String(status, err.Error())
}
