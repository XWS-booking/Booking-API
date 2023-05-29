package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type RateAccommodationDto struct {
	AccommodationId primitive.ObjectID `json:"accommodationId"`
	GuestId         primitive.ObjectID `json:"guestId"`
	Rating          int32              `json:"rating"`
}

type RateAccommodationHandler struct {
	ratingClientAddress      string
	reservationClientAddress string
}

func NewRateAccommodationHandler(ratingClientAddress string, reservationClientAddress string) Handler {
	return &RateAccommodationHandler{
		ratingClientAddress:      ratingClientAddress,
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *RateAccommodationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/rating/accommodation", handler.RateAccommodation)
	if err != nil {
		panic(err)
	}
}

func (handler *RateAccommodationHandler) RateAccommodation(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	//reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	var body RateAccommodationDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	fmt.Println(body.AccommodationId)
	res, err := ratingClient.RateAccommodation(context.TODO(), &gateway.RateAccommodationRequest{AccommodationId: body.AccommodationId.Hex(), GuestId: body.GuestId.Hex(), Rating: body.Rating})
	if err != nil {
		http.Error(w, "Failed rating accommodation!", http.StatusBadRequest)
		return
	}
	shared.Ok(&w, res)
}
