package api

import (
	"context"
	"encoding/json"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"github.com/golang/protobuf/ptypes"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type FindAllReservationsByAccommodationIdHandler struct {
	reservationClientAddress string
}

func NewFindAllReservationsByAccommodationIdHandler(authClientAddress, accommodationClientAddress, reservationClientAddress string) Handler {
	return &FindAllReservationsByAccommodationIdHandler{
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *FindAllReservationsByAccommodationIdHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/reservations/accommodation/{accommodationId}", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.FindAll))))
	if err != nil {
		panic(err)
	}
}

func (handler *FindAllReservationsByAccommodationIdHandler) FindAll(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["accommodationId"]
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	reservations, e := reservationClient.FindAllByAccommodationId(context.TODO(), &gateway.FindAllReservationsByAccommodationIdRequest{AccommodationId: id})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var reservationList []Reservation
	for _, r := range reservations.Reservations {
		startDate, _ := ptypes.Timestamp(r.StartDate)
		endDate, _ := ptypes.Timestamp(r.EndDate)
		reservationList = append(reservationList, Reservation{
			Id:            r.Id,
			Accommodation: Accommodation{},
			BuyerId:       r.BuyerId,
			StartDate:     startDate,
			EndDate:       endDate,
			Guests:        r.Guests,
			Status:        r.Status,
		})
	}

	response, err := json.Marshal(reservationList)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
