package api

import (
	"context"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type DeleteAccommodationRatingHandler struct {
	ratingClientAddress string
}

func NewDeleteAccommodationRatingHandler(ratingClientAddress string) Handler {
	return &DeleteAccommodationRatingHandler{
		ratingClientAddress: ratingClientAddress,
	}
}

func (handler *DeleteAccommodationRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/rating/accommodation/{id}", handler.DeleteRating)
	if err != nil {
		panic(err)
	}
}

func (handler *DeleteAccommodationRatingHandler) DeleteRating(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)

	res, err := ratingClient.DeleteAccommodationRating(context.TODO(), &gateway.DeleteAccommodationRatingRequest{Id: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	shared.Ok(&w, res)
}
