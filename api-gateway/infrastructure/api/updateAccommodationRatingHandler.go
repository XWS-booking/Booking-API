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

type UpdateAccommodationRatingDto struct {
	Id              string `json:"id"`
	Rating          int32  `json:"rating"`
	AccommodationId string `json:"accommodationId"`
	OldRating       int32  `json:"oldRating"`
}

type UpdateAccommodationRatingHandler struct {
	ratingClientAddress        string
	accommodationClientAddress string
	notificationClientAddress  string
}

func NewUpdateAccommodationRatingHandler(ratingClientAddress, accommodationClientAddress, notificationClientAddress string) Handler {
	return &UpdateAccommodationRatingHandler{
		ratingClientAddress:        ratingClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		notificationClientAddress:  notificationClientAddress,
	}
}

func (handler *UpdateAccommodationRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/rating/accommodation", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.UpdateRating))))
	if err != nil {
		panic(err)
	}
}

func (handler *UpdateAccommodationRatingHandler) UpdateRating(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	var body UpdateAccommodationRatingDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	res, err := ratingClient.UpdateAccommodationRating(context.TODO(), &gateway.UpdateAccommodationRatingRequest{Id: body.Id, Rating: body.Rating})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	accommodation, err := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: body.AccommodationId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.SendNotificationRequest{NotificationType: "guest_rated_accommodation", UserId: accommodation.OwnerId, Message: "Someone changed rating for '" + accommodation.Name + "' from " + strconv.Itoa(int(body.OldRating)) + " to " + strconv.Itoa(int(body.Rating))})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, res)
}
