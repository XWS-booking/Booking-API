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
	reservationClientAddress string
	authClientAddress        string
}

func NewCancelReservationHandler(reservationClientAddress, authClientAddress string) Handler {
	return &CancelReservationHandler{
		reservationClientAddress: reservationClientAddress,
		authClientAddress:        authClientAddress,
	}
}

func (handler *CancelReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/reservation/cancel", TokenValidationMiddleware(RolesMiddleware([]UserRole{1}, UserMiddleware(handler.Cancel))))
	if err != nil {
		panic(err)
	}
}

func (handler *CancelReservationHandler) Cancel(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	fmt.Println("Hit")
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
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
	shared.Ok(&w, res)
}
