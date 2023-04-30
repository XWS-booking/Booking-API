package main

import (
	"context"
	"gateway/infrastructure/api"
	"gateway/proto/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
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

	gwServer := &http.Server{
		Addr:    ":" + os.Getenv("GATEWAY_ADDRESS"),
		Handler: gwmux,
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
	authEndpoint := "auth_service:9000"
	accommodationEndpoint := "accomodation_service:9000"
	reservationEndpoint := "reservation_service:9000"

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

}
