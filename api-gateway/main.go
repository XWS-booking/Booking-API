package main

import (
	"context"
	"gateway/infrastructure/api"
	"gateway/proto/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {

	gwmux := runtime.NewServeMux()
	initHandlers(gwmux)
	handler := initCors(gwmux)
	gwServer := &http.Server{
		Addr:    ":" + os.Getenv("GATEWAY_ADDRESS"),
		Handler: handler,
	}

	go func() {
		if err := gwServer.ListenAndServe(); err != nil {
			log.Fatal("server error: ", err)
		}
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err := gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}

func initHandlers(gwmux *runtime.ServeMux) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authEndpoint := os.Getenv("AUTH_SERVICE_ADDRESS")
	accommodationEndpoint := os.Getenv("ACCOMODATION_SERVICE_ADDRESS")
	reservationEndpoint := os.Getenv("RESERVATION_SERVICE_ADDRESS")

	err := gateway.RegisterAuthServiceHandlerFromEndpoint(context.TODO(), gwmux, authEndpoint, opts)
	if err != nil {
		panic(err)
	}
	err = gateway.RegisterAccomodationServiceHandlerFromEndpoint(context.TODO(), gwmux, accommodationEndpoint, opts)
	if err != nil {
		panic(err)
	}
	err = gateway.RegisterReservationServiceHandlerFromEndpoint(context.TODO(), gwmux, reservationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	//init custom handlers
	searchAccommodationsHandler := api.NewSearchAccommodationHandler(accommodationEndpoint, reservationEndpoint)
	searchAccommodationsHandler.Init(gwmux)
	deleteProfileHandler := api.NewDeleteProfileHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	deleteProfileHandler.Init(gwmux)

	createAccomodationHandler := api.NewCreateAccomodationHandler(accommodationEndpoint, authEndpoint)
	createAccomodationHandler.Init(gwmux)

}

func initCors(gwmux *runtime.ServeMux) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(gwmux)
}
