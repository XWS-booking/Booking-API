package api

import (
	"context"
	"gateway/infrastructure/services"
	"gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/golang/protobuf/ptypes"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
	"time"
)

type AccommodationRatingDto struct {
	Id              string     `json:"id"`
	AccommodationId string     `json:"accommodationId"`
	Guest           model.User `json:"guest"`
	Rating          int32      `json:"rating"`
	Time            time.Time  `json:"time"`
}

type FindAllAccommodationRatingsHandler struct {
	ratingClientAddress string
	authClientAddress   string
}

func NewFindAllAccommodationRatingsHandler(ratingClientAddress string, authClientAddress string) Handler {
	return &FindAllAccommodationRatingsHandler{
		ratingClientAddress: ratingClientAddress,
		authClientAddress:   authClientAddress,
	}
}

func (handler *FindAllAccommodationRatingsHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/rating/accommodation/{accommodationId}", handler.FindAll)
	if err != nil {
		panic(err)
	}
}

func (handler *FindAllAccommodationRatingsHandler) FindAll(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	id := pathParams["accommodationId"]

	res, err := ratingClient.GetAllAccommodationRatings(context.TODO(), &gateway.GetAllAccommodationRatingsRequest{AccommodationId: id})
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	var ratingList []AccommodationRatingDto
	authClient := services.NewAuthClient(handler.authClientAddress)
	for _, r := range res.Ratings {
		guest, err := authClient.FindById(context.TODO(), &gateway.FindUserByIdRequest{Id: r.GuestId})
		if err != nil {
			shared.BadRequest(w, err.Error())
		}
		time, _ := ptypes.Timestamp(r.Time)
		ratingList = append(ratingList, AccommodationRatingDto{
			Id:              r.Id,
			AccommodationId: r.AccommodationId,
			Guest:           mapper.UserFromFindUserByIdResponse(guest),
			Rating:          r.Rating,
			Time:            time,
		})
	}
	shared.Ok(&w, ratingList)
}
