package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type UpdateAccommodationRatingDto struct {
	Id     string `json:"id"`
	Rating int32  `json:"rating"`
}

type UpdateAccommodationRatingHandler struct {
	ratingClientAddress string
}

func NewUpdateAccommodationRatingHandler(ratingClientAddress string) Handler {
	return &UpdateAccommodationRatingHandler{
		ratingClientAddress: ratingClientAddress,
	}
}

func (handler *UpdateAccommodationRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/rating/accommodation", handler.UpdateRating)
	if err != nil {
		panic(err)
	}
}

func (handler *UpdateAccommodationRatingHandler) UpdateRating(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)

	var body UpdateAccommodationRatingDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	res, err := ratingClient.UpdateAccommodationRating(context.TODO(), &gateway.UpdateAccommodationRatingRequest{Id: body.Id, Rating: body.Rating})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shared.Ok(&w, res)
}
