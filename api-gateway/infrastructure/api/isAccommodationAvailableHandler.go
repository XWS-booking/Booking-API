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

type IsAccommodationAvailableHandler struct {
	reservationClientAddress string
}

func NewIsAccommodationAvailableHandler(reservationClientAddress string) Handler {
	return &IsAccommodationAvailableHandler{
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *IsAccommodationAvailableHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/reservation/isAccommodationAvailable/{accommodationId}/{startDate}/{endDate}", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.IsAvailable))))
	if err != nil {
		panic(err)
	}
}

func (handler *IsAccommodationAvailableHandler) IsAvailable(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["accommodationId"]
	startDate, e := parseTimestamp(pathParams["startDate"])
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	endDate, e := parseTimestamp(pathParams["endDate"])
	if e != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	response, err := reservationClient.IsAccommodationAvailable(context.TODO(), &gateway.IsAccommodationAvailableRequest{AccommodationId: id, StartDate: startDate, EndDate: endDate})

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	shared.Ok(&w, response)
}
