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
	"strconv"
)

type UpdateRateHostDto struct {
	Id        string `json:"id"`
	Rating    int32  `json:"rating"`
	HostId    string `json:"hostId"`
	OldRating int32  `json:"oldRating"`
}

type UpdateHostRatingHandler struct {
	ratingClientAddress        string
	notificationClientAddress  string
	accommodationClientAddress string
	authClientAddress          string
	reservationClientAddress   string
}

func NewUpdateHostRatingHandler(accommodationClientAddress, authClientAddress, reservationClientAddress, ratingClientAddress, notificationClient string) Handler {
	return &UpdateHostRatingHandler{
		ratingClientAddress:        ratingClientAddress,
		notificationClientAddress:  notificationClient,
		accommodationClientAddress: accommodationClientAddress,
		authClientAddress:          authClientAddress,
		reservationClientAddress:   reservationClientAddress,
	}
}

func (handler *UpdateHostRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/rating/host", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.UpdateHostRate))))
	if err != nil {
		panic(err)
	}
}

func (handler *UpdateHostRatingHandler) UpdateHostRate(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	var body UpdateRateHostDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	res, err := ratingClient.UpdateHostRating(context.TODO(), &gateway.UpdateHostRatingRequest{Id: body.Id, Rating: body.Rating})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.SendNotificationRequest{NotificationType: "guest_rated_host", UserId: body.HostId, Message: "Someone changed your rating from " + strconv.Itoa(int(body.OldRating)) + " to " + strconv.Itoa(int(body.Rating))})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	hostDistinguishedChecker := NewIsHostDistinguishedFunc(handler.notificationClientAddress, handler.authClientAddress, handler.ratingClientAddress, handler.reservationClientAddress, handler.accommodationClientAddress)
	hostDistinguishedChecker.CheckIsHostDistinguishedFunc(body.HostId)
	shared.Ok(&w, res)
}
