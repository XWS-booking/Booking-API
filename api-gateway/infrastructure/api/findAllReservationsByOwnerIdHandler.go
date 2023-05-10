package api

import (
	"context"
	"encoding/json"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	ctx "github.com/gorilla/context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type FindAllReservationsByOwnerIdHandler struct {
	authClientAddress          string
	accommodationClientAddress string
	reservationClientAddress   string
}

func NewFindAllReservationsByOwnerIdHandler(authClientAddress, accommodationClientAddress, reservationClientAddress string) Handler {
	return &FindAllReservationsByOwnerIdHandler{
		authClientAddress:          authClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
	}
}

func (handler *FindAllReservationsByOwnerIdHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/reservations/owner", TokenValidationMiddleware(RolesMiddleware([]UserRole{1}, UserMiddleware(handler.FindAll))))
	if err != nil {
		panic(err)
	}
}

func (handler *FindAllReservationsByOwnerIdHandler) FindAll(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := ctx.Get(r, "id").(string)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	accommodations, e := accommodationClient.FindAllAccommodationIdsByOwnerId(context.TODO(), &gateway.FindAllAccommodationIdsByOwnerIdRequest{OwnerId: id})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	var reservationsWithAccommodation []ReservationWithCancellation
	for _, accommId := range accommodations.Ids {
		accommodation, e := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: accommId})
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reservations, e := reservationClient.FindAllByAccommodationId(context.TODO(), &gateway.FindAllReservationsByAccommodationIdRequest{AccommodationId: accommId})
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		for _, reservation := range reservations.Reservations {
			numberOfCancellation, e := reservationClient.FindNumberOfBuyersCancellations(context.TODO(), &gateway.NumberOfCancellationRequest{BuyerId: reservation.BuyerId})
			if e != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			reservationsWithAccommodation = append(reservationsWithAccommodation, ReservationWithCancellation{
				Id:                   reservation.Id,
				Accommodation:        mapper.AccommodationFromAccomodationResponse(accommodation, User{}),
				BuyerId:              reservation.BuyerId,
				StartDate:            reservation.StartDate.AsTime(),
				EndDate:              reservation.EndDate.AsTime(),
				Guests:               reservation.Guests,
				Status:               reservation.Status,
				NumberOfCancellation: numberOfCancellation.CancellationNumber,
			})
		}
	}
	response, err := json.Marshal(reservationsWithAccommodation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
