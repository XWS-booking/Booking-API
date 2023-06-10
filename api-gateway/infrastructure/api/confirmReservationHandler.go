package api

import (
	"context"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type ConfirmReservationDto struct {
	ReservationId primitive.ObjectID `json:"reservationId"`
}

type ConfirmReservationHandler struct {
	reservationClientAddress  string
	notificationClientAddress string
}

func NewConfirmReservationHandler(reservationClientAddress, notificationClientAddress string) Handler {
	return &ConfirmReservationHandler{
		reservationClientAddress:  reservationClientAddress,
		notificationClientAddress: notificationClientAddress,
	}
}

func (handler *ConfirmReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/reservation/confirm/{id}", TokenValidationMiddleware(RolesMiddleware([]UserRole{1}, handler.Confirm)))
	if err != nil {
		panic(err)
	}
}

func (handler *ConfirmReservationHandler) Confirm(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	id, e := pathParams["id"]
	if !e {
		shared.BadRequest(w, "Error with data!")
		return
	}
	res, err := reservationClient.Confirm(context.TODO(), &gateway.ReservationId{Id: id})
	if err != nil {
		shared.BadRequest(w, "Error when confirming reservation!")
		return
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.NotificationRequest{UserId: res.BuyerId, Message: "Host confirmed your reservation!"})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, res)
}
