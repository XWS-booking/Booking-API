package api

import (
	"context"
	"encoding/json"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	"gateway/model"
	. "gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	"github.com/golang/protobuf/ptypes"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type FindAllReservationsByBuyerIdHandler struct {
	authClientAddress          string
	accommodationClientAddress string
	reservationClientAddress   string
	ratingClientAddress        string
}

func NewFindAllReservationsByBuyerIdHandler(authClientAddress, accommodationClientAddress, reservationClientAddress, ratingClientAddress string) Handler {
	return &FindAllReservationsByBuyerIdHandler{
		authClientAddress:          authClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
		ratingClientAddress:        ratingClientAddress,
	}
}

func (handler *FindAllReservationsByBuyerIdHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/reservations/buyer", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.FindAll))))
	if err != nil {
		panic(err)
	}
}

func (handler *FindAllReservationsByBuyerIdHandler) FindAll(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	token := r.Header["Authorization"][0]

	authClient := services.NewAuthClient(handler.authClientAddress)
	user, e := authClient.GetUser(context.TODO(), &gateway.GetUserRequest{Token: token})
	if e != nil {
		panic(e)
	}
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	reservations, e := reservationClient.FindAllByBuyerId(context.TODO(), &gateway.FindAllReservationsByBuyerIdRequest{BuyerId: user.Id})
	if e != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var reservationsWithAccommodation []model.Reservation
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	for _, r := range reservations.Reservations {
		accommodation, e := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: r.AccommodationId})
		if e != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		var rating Rating
		if r.AccommodationRatingId == "000000000000000000000000" {
			rating = Rating{}
		} else {
			resp, _ := ratingClient.FindAccommodationRatingById(context.TODO(), &gateway.FindAccommodationRatingByIdRequest{Id: r.AccommodationRatingId})
			rating = Rating{Id: resp.Id, AccommodationId: resp.AccommodationId, GuestId: resp.GuestId, Rating: resp.Rating}
		}
		startDate, _ := ptypes.Timestamp(r.StartDate)
		endDate, _ := ptypes.Timestamp(r.EndDate)
		reservationsWithAccommodation = append(reservationsWithAccommodation, model.Reservation{
			Id:                  r.Id,
			Accommodation:       mapper.AccommodationFromAccomodationResponse(accommodation, model.User{}, 0),
			BuyerId:             r.BuyerId,
			StartDate:           startDate,
			EndDate:             endDate,
			Guests:              r.Guests,
			Status:              r.Status,
			AccommodationRating: rating,
		})
	}
	response, err := json.Marshal(reservationsWithAccommodation)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
