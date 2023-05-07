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

type RejectReservationDto struct {
	ReservationId primitive.ObjectID `json:"reservationId"`
}

type RejectReservationHandler struct {
	reservationClientAddress string
}

func NewRejectReservationHandler(reservationClientAddress string) Handler {
	return &RejectReservationHandler{
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *RejectReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/reservation/reject", TokenValidationMiddleware(RolesMiddleware([]UserRole{1}, handler.Reject)))
	if err != nil {
		panic(err)
	}
}

func (handler *RejectReservationHandler) Reject(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	var body RejectReservationDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	res, err := reservationClient.Reject(context.TODO(), &gateway.ReservationId{Id: body.ReservationId.Hex()})
	if err != nil {
		shared.BadRequest(w, "Error when rejecting reservation!")
		return
	}
	shared.Ok(&w, res)
}
