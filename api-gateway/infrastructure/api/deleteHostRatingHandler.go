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

type DeleteRateHostDto struct {
	Id string `json:"id"`
}

type DeleteHostRatingHandler struct {
	ratingClientAddress string
}

func NewDeleteHostRatingHandler(ratingClientAddress string) Handler {
	return &DeleteHostRatingHandler{
		ratingClientAddress: ratingClientAddress,
	}
}

func (handler *DeleteHostRatingHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/rating/{id}/host", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.DeleteHostRate))))
	if err != nil {
		panic(err)
	}
}

func (handler *DeleteHostRatingHandler) DeleteHostRate(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	ratingId := pathParams["id"]

	res, err := ratingClient.DeleteHostRating(context.TODO(), &gateway.DeleteHostRatingRequest{Id: ratingId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	shared.Ok(&w, res)
}
