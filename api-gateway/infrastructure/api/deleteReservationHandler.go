package api

import (
	"context"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type DeleteReservationHandler struct {
	reservationClientAddress string
}

func NewDeleteReservationHandler(reservationClientAddress string) Handler {
	return &DeleteReservationHandler{
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *DeleteReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/reservation/{id}", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.Delete))))
	if err != nil {
		panic(err)
	}
}

func (handler *DeleteReservationHandler) Delete(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]

	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	response, err := reservationClient.Delete(context.TODO(), &gateway.ReservationId{Id: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.Ok(&w, response)
}
