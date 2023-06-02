package api

import (
	"context"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type HostRatingsDto struct {
	AccommodationId string `json:"accommodationId"`
	GuestId         string `json:"guestId"`
	Rating          int32  `json:"rating"`
	ReservationId   string `json:"reservationId"`
}

type GetHostRatingsHandler struct {
	ratingClientAddress string
	authClientAddress   string
}

func NewGetHostRatingsHandler(ratingClientAddress string, authClientAddress string) Handler {
	return &GetHostRatingsHandler{
		ratingClientAddress: ratingClientAddress,
		authClientAddress:   authClientAddress,
	}
}

func (handler *GetHostRatingsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/rating/{id}/host", handler.GetHostRatings)
	if err != nil {
		panic(err)
	}
}

func (handler *GetHostRatingsHandler) GetHostRatings(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	hostId := pathParams["id"]

	res, err := ratingClient.GetHostRatings(context.TODO(), &gateway.GetHostRatingsRequest{HostId: hostId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	res2, err2 := authClient.GetHostRatingWithGuestInfo(context.TODO(), &gateway.GetHostRatingWithGuestInfoRequest{Ratings: res.Ratings})
	res2.AverageRate = res.AverageRate

	if err2 != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, res2)
}
