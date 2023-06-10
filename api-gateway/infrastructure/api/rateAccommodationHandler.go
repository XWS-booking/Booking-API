package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type RateAccommodationDto struct {
	AccommodationId string `json:"accommodationId"`
	GuestId         string `json:"guestId"`
	Rating          int32  `json:"rating"`
	ReservationId   string `json:"reservationId"`
}

type RateAccommodationHandler struct {
	ratingClientAddress        string
	reservationClientAddress   string
	notificationClientAddress  string
	accommodationClientAddress string
}

func NewRateAccommodationHandler(ratingClientAddress, reservationClientAddress, notificationClientAddress, accommodationClientAddress string) Handler {
	return &RateAccommodationHandler{
		ratingClientAddress:        ratingClientAddress,
		reservationClientAddress:   reservationClientAddress,
		notificationClientAddress:  notificationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
	}
}

func (handler *RateAccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/rating/accommodation", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.RateAccommodation))))
	if err != nil {
		panic(err)
	}
}

func (handler *RateAccommodationHandler) RateAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	var body RateAccommodationDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	fmt.Println(body.AccommodationId)
	res, err := ratingClient.RateAccommodation(context.TODO(), &gateway.RateAccommodationRequest{AccommodationId: body.AccommodationId, GuestId: body.GuestId, Rating: body.Rating})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	_, err2 := reservationClient.UpdateReservationRating(context.TODO(), &gateway.UpdateReservationRatingRequest{Id: body.ReservationId, AccommodationRatingId: res.Id})
	if err2 != nil {
		shared.BadRequest(w, err2.Error())
		return
	}
	accommodation, err := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: body.AccommodationId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.SendNotificationRequest{NotificationType: "guest_rated_accommodation", UserId: accommodation.OwnerId, Message: "Someone rated your accommodation '" + accommodation.Name + "'"})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, res)
}
