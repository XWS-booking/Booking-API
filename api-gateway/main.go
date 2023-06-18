package main

import (
	"context"
	"gateway/infrastructure/api"
	metrics "gateway/metrics"
	"gateway/proto/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
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
	reg := prometheus.NewRegistry()
	metHttp := metrics.NewMetricsHttp(reg)
	metHttpSuccess := metrics.NewMetricsHttpSuccess(reg)
	metHttpError := metrics.NewMetricsHttpError(reg)
	gwmux := runtime.NewServeMux()
	initHandlers(gwmux, metHttp, metHttpSuccess, metHttpError)
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

	promHandler := promhttp.HandlerFor(reg, promhttp.HandlerOpts{})
	http.Handle("/metrics", promHandler)
	go func() {
		http.ListenAndServe(":8086", nil)
	}()

	stopCh := make(chan os.Signal)
	signal.Notify(stopCh, syscall.SIGTERM)

	<-stopCh

	if err := gwServer.Close(); err != nil {
		log.Fatalln("error while stopping server: ", err)
	}
}

func initHandlers(gwmux *runtime.ServeMux, metHttp, metHttpSuccess, metHttpError *metrics.MetricsHttp) {
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	authEndpoint := os.Getenv("AUTH_SERVICE_ADDRESS")
	accommodationEndpoint := os.Getenv("ACCOMODATION_SERVICE_ADDRESS")
	reservationEndpoint := os.Getenv("RESERVATION_SERVICE_ADDRESS")
	ratingEndpoint := os.Getenv("RATING_SERVICE_ADDRESS")
	notificationEndpoint := os.Getenv("NOTIFICATION_SERVICE_ADDRESS")
	recommendationEndpoint := os.Getenv("RECOMMENDATION_SERVICE_ADDRESS")
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
	searchAccommodationsHandler := api.NewSearchAccommodationHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, accommodationEndpoint, reservationEndpoint, ratingEndpoint)
	searchAccommodationsHandler.Init(gwmux)
	deleteProfileHandler := api.NewDeleteProfileHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, accommodationEndpoint, reservationEndpoint)
	deleteProfileHandler.Init(gwmux)
	createAccommodationHandler := api.NewCreateAccomodationHandler(metHttp, metHttpSuccess, metHttpError, accommodationEndpoint, authEndpoint)
	createAccommodationHandler.Init(gwmux)
	cancelReservationHandler := api.NewCancelReservationHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, authEndpoint, reservationEndpoint, accommodationEndpoint, notificationEndpoint)
	cancelReservationHandler.Init(gwmux)
	createReservationHandler := api.NewCreateReservationHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, reservationEndpoint, authEndpoint, accommodationEndpoint, notificationEndpoint)
	createReservationHandler.Init(gwmux)
	confirmReservationHandler := api.NewConfirmReservationHandler(metHttp, metHttpSuccess, metHttpError, reservationEndpoint, notificationEndpoint)
	confirmReservationHandler.Init(gwmux)
	rejectReservationHandler := api.NewRejectReservationHandler(metHttp, metHttpSuccess, metHttpError, reservationEndpoint, notificationEndpoint)
	rejectReservationHandler.Init(gwmux)
	findAllReservationsByOwnerIdHandler := api.NewFindAllReservationsByOwnerIdHandler(metHttp, metHttpSuccess, metHttpError, accommodationEndpoint, reservationEndpoint)
	findAllReservationsByOwnerIdHandler.Init(gwmux)
	findAllReservationsByBuyerIdHandler := api.NewFindAllReservationsByBuyerIdHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, accommodationEndpoint, reservationEndpoint, ratingEndpoint)
	findAllReservationsByBuyerIdHandler.Init(gwmux)
	updatePersonalInfoHandler := api.NewUpdatePersonalInfoHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint)
	updatePersonalInfoHandler.Init(gwmux)
	changePasswordHandler := api.NewChangePasswordHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint)
	changePasswordHandler.Init(gwmux)
	deleteReservationHandler := api.NewDeleteReservationHandler(metHttp, metHttpSuccess, metHttpError, notificationEndpoint, accommodationEndpoint, authEndpoint, reservationEndpoint, ratingEndpoint)
	deleteReservationHandler.Init(gwmux)
	isAccommodationAvailableHandler := api.NewIsAccommodationAvailableHandler(metHttp, metHttpSuccess, metHttpError, reservationEndpoint)
	isAccommodationAvailableHandler.Init(gwmux)
	updatePricingHandler := api.NewUpdatePricingHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, accommodationEndpoint, reservationEndpoint)
	updatePricingHandler.Init(gwmux)
	findAllReservationsByAccommodationIdHandler := api.NewFindAllReservationsByAccommodationIdHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, accommodationEndpoint, reservationEndpoint)
	findAllReservationsByAccommodationIdHandler.Init(gwmux)
	getBookingPriceHandler := api.NewGetBookingPriceHandler(metHttp, metHttpSuccess, metHttpError, accommodationEndpoint)
	getBookingPriceHandler.Init(gwmux)
	rateAccommodationHandler := api.NewRateAccommodationHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, reservationEndpoint, notificationEndpoint, accommodationEndpoint)
	rateAccommodationHandler.Init(gwmux)
	deleteAccommodationRatingHandler := api.NewDeleteAccommodationRatingHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, reservationEndpoint)
	deleteAccommodationRatingHandler.Init(gwmux)
	updateAccommodationRatingHandler := api.NewUpdateAccommodationRatingHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, accommodationEndpoint, notificationEndpoint)
	updateAccommodationRatingHandler.Init(gwmux)
	findAllAccommodationRatingsHandler := api.NewFindAllAccommodationRatingsHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, authEndpoint)
	findAllAccommodationRatingsHandler.Init(gwmux)
	rateHostHandler := api.NewRateHostHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, ratingEndpoint, reservationEndpoint, accommodationEndpoint, notificationEndpoint)
	rateHostHandler.Init(gwmux)
	updateHostRateHandler := api.NewUpdateHostRatingHandler(metHttp, metHttpSuccess, metHttpError, accommodationEndpoint, authEndpoint, reservationEndpoint, ratingEndpoint, notificationEndpoint)
	updateHostRateHandler.Init(gwmux)
	deleteHostRatingHandler := api.NewDeleteHostRatingHandler(metHttp, metHttpSuccess, metHttpError, notificationEndpoint, accommodationEndpoint, authEndpoint, reservationEndpoint, ratingEndpoint)
	deleteHostRatingHandler.Init(gwmux)
	getHostRatingsHandler := api.NewGetHostRatingsHandler(metHttp, metHttpSuccess, metHttpError, ratingEndpoint, authEndpoint)
	getHostRatingsHandler.Init(gwmux)
	registerUserHandler := api.NewRegisterUserHandler(metHttp, metHttpSuccess, metHttpError, authEndpoint, notificationEndpoint)
	registerUserHandler.Init(gwmux)
	findNotificationPreferencesByUserId := api.NewFindNotificationPreferencesByUserHandler(metHttp, metHttpSuccess, metHttpError, notificationEndpoint)
	findNotificationPreferencesByUserId.Init(gwmux)
	updateNotificationPreferencesHandler := api.NewUpdateNotificationPreferencesHandler(metHttp, metHttpSuccess, metHttpError, notificationEndpoint)
	updateNotificationPreferencesHandler.Init(gwmux)
	getRecommendedAccommodationsHandler := api.NewRecommendedAccommodationsHandler(metHttp, metHttpSuccess, metHttpError, accommodationEndpoint, recommendationEndpoint, authEndpoint, ratingEndpoint)
	getRecommendedAccommodationsHandler.Init(gwmux)
}

func initCors(gwmux *runtime.ServeMux) http.Handler {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "PATCH"},
		AllowedHeaders:   []string{"Content-Type", "Authorization", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
		MaxAge:           300,
	}).Handler(gwmux)
}
