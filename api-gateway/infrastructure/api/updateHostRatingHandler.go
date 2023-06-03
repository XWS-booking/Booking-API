package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type UpdateRateHostDto struct {
	Id     string `json:"id"`
	Rating int32  `json:"rating"`
}

type UpdateHostRatingHandler struct {
	ratingClientAddress string
}

func NewUpdateHostRatingHandler(ratingClientAddress string) Handler {
	return &UpdateHostRatingHandler{
		ratingClientAddress: ratingClientAddress,
	}
}

func (handler *UpdateHostRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("PATCH", "/api/rating/host", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.UpdateHostRate))))
	if err != nil {
		panic(err)
	}
}

func (handler *UpdateHostRatingHandler) UpdateHostRate(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	var body UpdateRateHostDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}

	res, err := ratingClient.UpdateHostRating(context.TODO(), &gateway.UpdateHostRatingRequest{Id: body.Id, Rating: body.Rating})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	shared.Ok(&w, res)
}
