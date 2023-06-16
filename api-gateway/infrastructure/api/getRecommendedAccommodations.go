package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	"gateway/model"
	. "gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	"gateway/shared"
	ctx "github.com/gorilla/context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type GetRecommendedAccommodationsHandler struct {
	accommodationClientAddress  string
	recommendationClientAddress string
	authClientAddress           string
	ratingClientClientAddress   string
}

func NewRecommendedAccommodationsHandler(accommodationClientAddress string, recommendationClientAddress string, authClientAddress string, ratingClientAddress string) Handler {
	return &GetRecommendedAccommodationsHandler{
		recommendationClientAddress: recommendationClientAddress,
		accommodationClientAddress:  accommodationClientAddress,
		authClientAddress:           authClientAddress,
		ratingClientClientAddress:   ratingClientAddress,
	}
}

func (handler *GetRecommendedAccommodationsHandler) Init(mux *runtime.ServeMux) {

	err := mux.HandlePath("GET", "/api/accomodation/recommended", TokenValidationMiddleware(RolesMiddleware([]UserRole{HOST, GUEST}, UserMiddleware(handler.GetRecommendedAccommodations))))
	if err != nil {
		panic(err)
	}
}

func (handler *GetRecommendedAccommodationsHandler) GetRecommendedAccommodations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	recommendationClient := services.NewRecommendationClient(handler.recommendationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	authClient := services.NewAuthClient(handler.authClientAddress)
	ratingClient := services.NewRatingClient(handler.ratingClientClientAddress)

	userId := ctx.Get(r, "id").(string)
	fmt.Println("User id is ", userId)

	resp, err := recommendationClient.GetRecommendedAccommodations(context.TODO(), &gateway.RecommendationRequest{UserId: userId})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	accommodations, err := accommodationClient.PopulateRecommended(context.TODO(), &gateway.PopulateRecommendedRequest{
		Ids: resp.Accommodations,
	})

	data := make([]model.Accommodation, 0)

	for _, single := range accommodations.Accommodations {
		owner, err := authClient.FindById(context.TODO(), &gateway.FindUserByIdRequest{Id: single.OwnerId})
		if err != nil {
			shared.BadRequest(w, "Something wrong with capturing owner")
			return
		}
		averageRating, err := ratingClient.GetAverageAccommodationRating(context.TODO(), &gateway.GetAverageAccommodationRatingRequest{AccommodationId: single.Id})
		if err != nil {
			shared.BadRequest(w, "Something wrong with capturing rating")
			return
		}
		data = append(data, mapper.AccommodationFromAccomodationResponse(single, mapper.UserFromFindUserByIdResponse(owner), averageRating.Rating))
	}

	if err != nil {
		shared.BadRequest(w, "Something wrong with capturing recommended accommodations!")
		return
	}

	shared.Ok(&w, data)
}
