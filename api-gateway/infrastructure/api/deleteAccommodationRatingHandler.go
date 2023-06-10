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

type DeleteAccommodationRatingHandler struct {
	ratingClientAddress      string
	reservationClientAddress string
}

func NewDeleteAccommodationRatingHandler(ratingClientAddress, reservationClientAddress string) Handler {
	return &DeleteAccommodationRatingHandler{
		ratingClientAddress:      ratingClientAddress,
		reservationClientAddress: reservationClientAddress,
	}
}

func (handler *DeleteAccommodationRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/rating/accommodation/{id}/{reservationId}", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.DeleteRating))))
	if err != nil {
		panic(err)
	}
}

func (handler *DeleteAccommodationRatingHandler) DeleteRating(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingId := pathParams["id"]
	reservationId := pathParams["reservationId"]

	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)

	_, err := reservationClient.UpdateReservationRating(context.TODO(), &gateway.UpdateReservationRatingRequest{Id: reservationId, AccommodationRatingId: "000000000000000000000000"})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	res, err := ratingClient.DeleteAccommodationRating(context.TODO(), &gateway.DeleteAccommodationRatingRequest{Id: ratingId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, res)
}
