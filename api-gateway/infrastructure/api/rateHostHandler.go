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

type RateHostDto struct {
	HostId  string `json:"hostId"`
	GuestId string `json:"guestId"`
	Rating  int32  `json:"rating"`
}

type RateHostHandler struct {
	ratingClientAddress        string
	reservationClientAddress   string
	accommodationClientAddress string
	notificationClientAddress  string
}

func NewRateHostHandler(ratingClientAddress string, reservationClientAddress string, accommodationClientAddress string, notificationClientAddress string) Handler {
	return &RateHostHandler{
		ratingClientAddress:        ratingClientAddress,
		reservationClientAddress:   reservationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		notificationClientAddress:  notificationClientAddress,
	}
}

func (handler *RateHostHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/rating/host", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.RateHost))))
	if err != nil {
		panic(err)
	}
}

func (handler *RateHostHandler) RateHost(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	var body RateHostDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	res, err := accommodationClient.FindAllAccommodationIdsByOwnerId(context.TODO(), &gateway.FindAllAccommodationIdsByOwnerIdRequest{OwnerId: body.HostId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	res2, err2 := reservationClient.CheckIfGuestHasReservationInAccommodations(context.TODO(), &gateway.CheckIfGuestHasReservationInAccommodationsRequest{GuestId: body.GuestId, AccommodationIds: res.Ids})
	if err2 != nil {
		fmt.Println(err2)
		shared.BadRequest(w, err2.Error())
		return
	}

	if !res2.Res {
		shared.BadRequest(w, "You can't rate this host since you don't have reservation at this host in the past!")
		return
	}

	res3, err3 := ratingClient.RateHost(context.TODO(), &gateway.RateHostRequest{HostId: body.HostId, Rating: body.Rating, GuestId: body.GuestId})

	if err3 != nil {
		http.Error(w, fmt.Sprintf("Unsuccessful host rating!", err.Error()), http.StatusBadRequest)
		return
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.NotificationRequest{UserId: body.HostId, Message: "Someone rated you!"})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, res3)
}
