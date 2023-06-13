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
	ratingClientAddress        string
	accommodationClientAddress string
	authClientAddress          string
	reservationClientAddress   string
	notificationClientAddress  string
}

func NewDeleteReservationHandler(notificationClientAddress, accommodationClientAddress, authClientAddress, reservationClientAddress, ratingClientAddress string) Handler {
	return &DeleteReservationHandler{
		ratingClientAddress:        ratingClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		authClientAddress:          authClientAddress,
		reservationClientAddress:   reservationClientAddress,
		notificationClientAddress:  notificationClientAddress,
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
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	response, err := reservationClient.Delete(context.TODO(), &gateway.ReservationId{Id: id})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	resp, err := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: response.AccommodationId})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	hostDistinguishedChecker := NewIsHostDistinguishedFunc(handler.notificationClientAddress, handler.authClientAddress, handler.ratingClientAddress, handler.reservationClientAddress, handler.accommodationClientAddress)
	hostDistinguishedChecker.CheckIsHostDistinguishedFunc(resp.OwnerId)
	shared.Ok(&w, response)
}
