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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type CancelReservationDto struct {
	Token         string             `json:"token"`
	ReservationId primitive.ObjectID `json:"reservationId"`
}

type CancelReservationHandler struct {
	reservationClientAddress   string
	accommodationClientAddress string
	notificationClientAddress  string
	authClientAddress          string
	ratingClientAddress        string
}

func NewCancelReservationHandler(ratingClientAddress, authClientAddress, reservationClientAddress, accommodationClientAddress, notificationClientAddress string) Handler {
	return &CancelReservationHandler{
		reservationClientAddress:   reservationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		notificationClientAddress:  notificationClientAddress,
		authClientAddress:          authClientAddress,
		ratingClientAddress:        ratingClientAddress,
	}
}

func (handler *CancelReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/reservation/cancel", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.Cancel))))
	if err != nil {
		panic(err)
	}
}

func (handler *CancelReservationHandler) Cancel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	var body CancelReservationDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	res, err := reservationClient.CancelReservation(context.TODO(), &gateway.CancelReservationRequest{Token: "", ReservationId: body.ReservationId.Hex()})
	if err != nil {
		shared.BadRequest(w, "Error when canceling reservation!")
		return
	}

	accommodation, err := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: res.AccommodationId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.SendNotificationRequest{NotificationType: "guest_canceled_reservation", UserId: accommodation.OwnerId, Message: "Someone canceled reservation in'" + accommodation.Name + "'"})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	hostDistinguishedChecker := NewIsHostDistinguishedFunc(handler.notificationClientAddress, handler.authClientAddress, handler.ratingClientAddress, handler.reservationClientAddress, handler.accommodationClientAddress)
	hostDistinguishedChecker.CheckIsHostDistinguishedFunc(accommodation.OwnerId)
	shared.Ok(&w, res)
}
