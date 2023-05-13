package main

import (
	"context"
	"fmt"
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
	fmt.Println(err)
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
	searchAccommodationsHandler := api.NewSearchAccommodationHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	searchAccommodationsHandler.Init(gwmux)
	deleteProfileHandler := api.NewDeleteProfileHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	deleteProfileHandler.Init(gwmux)
	createAccommodationHandler := api.NewCreateAccomodationHandler(accommodationEndpoint, authEndpoint)
	createAccommodationHandler.Init(gwmux)
	cancelReservationHandler := api.NewCancelReservationHandler(reservationEndpoint, authEndpoint)
	cancelReservationHandler.Init(gwmux)
	createReservationHandler := api.NewCreateReservationHandler(reservationEndpoint, authEndpoint, accommodationEndpoint)
	createReservationHandler.Init(gwmux)
	confirmReservationHandler := api.NewConfirmReservationHandler(reservationEndpoint)
	confirmReservationHandler.Init(gwmux)
	rejectReservationHandler := api.NewRejectReservationHandler(reservationEndpoint)
	rejectReservationHandler.Init(gwmux)
	findAllReservationsByOwnerIdHandler := api.NewFindAllReservationsByOwnerIdHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	findAllReservationsByOwnerIdHandler.Init(gwmux)
	findAllReservationsByBuyerIdHandler := api.NewFindAllReservationsByBuyerIdHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	findAllReservationsByBuyerIdHandler.Init(gwmux)
	updatePersonalInfoHandler := api.NewUpdatePersonalInfoHandler(authEndpoint)
	updatePersonalInfoHandler.Init(gwmux)
	changePasswordHadler := api.NewChangePasswordHandler(authEndpoint)
	changePasswordHadler.Init(gwmux)
	deleteReservationHandler := api.NewDeleteReservationHandler(reservationEndpoint)
	deleteReservationHandler.Init(gwmux)
	isAccommodationAvailableHandler := api.NewIsAccommodationAvailableHandler(reservationEndpoint)
	isAccommodationAvailableHandler.Init(gwmux)
	updatePricingHandler := api.NewUpdatePricingHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	updatePricingHandler.Init(gwmux)
}

func initCors(gwmux *runtime.ServeMux) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(gwmux)
}
