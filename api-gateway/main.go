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
	ratingEndpoint := os.Getenv("RATING_SERVICE_ADDRESS")
	notificationEndpoint := os.Getenv("NOTIFICATION_SERVICE_ADDRESS")
	fmt.Println(notificationEndpoint)
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
	err = gateway.RegisterRatingServiceHandlerFromEndpoint(context.TODO(), gwmux, ratingEndpoint, opts)
	if err != nil {
		panic(err)
	}
	err = gateway.RegisterNotificationServiceHandlerFromEndpoint(context.TODO(), gwmux, notificationEndpoint, opts)
	if err != nil {
		panic(err)
	}

	//init custom handlers
	searchAccommodationsHandler := api.NewSearchAccommodationHandler(authEndpoint, accommodationEndpoint, reservationEndpoint, ratingEndpoint)
	searchAccommodationsHandler.Init(gwmux)
	deleteProfileHandler := api.NewDeleteProfileHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	deleteProfileHandler.Init(gwmux)
	createAccommodationHandler := api.NewCreateAccomodationHandler(accommodationEndpoint, authEndpoint)
	createAccommodationHandler.Init(gwmux)
	cancelReservationHandler := api.NewCancelReservationHandler(reservationEndpoint, accommodationEndpoint, notificationEndpoint)
	cancelReservationHandler.Init(gwmux)
	createReservationHandler := api.NewCreateReservationHandler(reservationEndpoint, authEndpoint, accommodationEndpoint, notificationEndpoint)
	createReservationHandler.Init(gwmux)
	confirmReservationHandler := api.NewConfirmReservationHandler(reservationEndpoint, notificationEndpoint)
	confirmReservationHandler.Init(gwmux)
	rejectReservationHandler := api.NewRejectReservationHandler(reservationEndpoint, notificationEndpoint)
	rejectReservationHandler.Init(gwmux)
	findAllReservationsByOwnerIdHandler := api.NewFindAllReservationsByOwnerIdHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	findAllReservationsByOwnerIdHandler.Init(gwmux)
	findAllReservationsByBuyerIdHandler := api.NewFindAllReservationsByBuyerIdHandler(authEndpoint, accommodationEndpoint, reservationEndpoint, ratingEndpoint)
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
	findAllReservationsByAccommodationIdHandler := api.NewFindAllReservationsByAccommodationIdHandler(authEndpoint, accommodationEndpoint, reservationEndpoint)
	findAllReservationsByAccommodationIdHandler.Init(gwmux)
	getBookingPriceHandler := api.NewGetBookingPriceHandler(accommodationEndpoint)
	getBookingPriceHandler.Init(gwmux)
	rateAccommodationHandler := api.NewRateAccommodationHandler(ratingEndpoint, reservationEndpoint, notificationEndpoint, accommodationEndpoint)
	rateAccommodationHandler.Init(gwmux)
	deleteAccommodationRatingHandler := api.NewDeleteAccommodationRatingHandler(ratingEndpoint, reservationEndpoint)
	deleteAccommodationRatingHandler.Init(gwmux)
	updateAccommodationRatingHandler := api.NewUpdateAccommodationRatingHandler(ratingEndpoint)
	updateAccommodationRatingHandler.Init(gwmux)
	findAllAccommodationRatingsHandler := api.NewFindAllAccommodationRatingsHandler(ratingEndpoint, authEndpoint)
	findAllAccommodationRatingsHandler.Init(gwmux)
	rateHostHandler := api.NewRateHostHandler(ratingEndpoint, reservationEndpoint, accommodationEndpoint, notificationEndpoint)
	rateHostHandler.Init(gwmux)
	updateHostRateHandler := api.NewUpdateHostRatingHandler(ratingEndpoint)
	updateHostRateHandler.Init(gwmux)
	deleteHostRatingHandler := api.NewDeleteHostRatingHandler(ratingEndpoint)
	deleteHostRatingHandler.Init(gwmux)
	getHostRatingsHandler := api.NewGetHostRatingsHandler(ratingEndpoint, authEndpoint)
	getHostRatingsHandler.Init(gwmux)
	registerUserHandler := api.NewRegisterUserHandler(authEndpoint, notificationEndpoint)
	registerUserHandler.Init(gwmux)
	findNotificationPreferencesByUserId := api.NewFindNotificationPreferencesByUserHandler(notificationEndpoint)
	findNotificationPreferencesByUserId.Init(gwmux)
	updateNotificationPreferencesHandler := api.NewUpdateNotificationPreferencesHandler(notificationEndpoint)
	updateNotificationPreferencesHandler.Init(gwmux)
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
