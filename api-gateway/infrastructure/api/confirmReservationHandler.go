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

type ConfirmReservationDto struct {
	ReservationId primitive.ObjectID `json:"reservationId"`
}

type ConfirmReservationHandler struct {
	reservationClientAddress string
}

func NewConfirmReservationHandler(reservationClientAddress string) Handler {
	return &ConfirmReservationHandler{
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *ConfirmReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/reservation/confirm", TokenValidationMiddleware(RolesMiddleware([]UserRole{1}, handler.Confirm)))
	if err != nil {
		panic(err)
	}
}

func (handler *ConfirmReservationHandler) Confirm(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	var body ConfirmReservationDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	res, err := reservationClient.Confirm(context.TODO(), &gateway.ReservationId{Id: body.ReservationId.Hex()})
	if err != nil {
		shared.BadRequest(w, "Error when confirming reservation!")
		return
	}
	shared.Ok(&w, res)
}
