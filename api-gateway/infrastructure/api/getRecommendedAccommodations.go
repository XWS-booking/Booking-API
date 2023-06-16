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

type GetRecommendedAccommodationsHandler struct {
	accommodationClientAddress  string
	recommendationClientAddress string
}

func NewRecommendedAccommodationsHandler(accommodationClientAddress string, recommendationClientAddress string) Handler {
	return &GetRecommendedAccommodationsHandler{
		recommendationClientAddress: recommendationClientAddress,
		accommodationClientAddress:  accommodationClientAddress,
	}
}

func (handler *GetRecommendedAccommodationsHandler) Init(mux *runtime.ServeMux) {

	err := mux.HandlePath("GET", "/api/accomodation/recommended", TokenValidationMiddleware(RolesMiddleware([]UserRole{HOST}, UserMiddleware(handler.GetRecommendedAccommodations))))
	if err != nil {
		panic(err)
	}
}

func (handler *GetRecommendedAccommodationsHandler) GetRecommendedAccommodations(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	recommendationClient := services.NewRecommendationClient(handler.recommendationClientAddress)
	fmt.Println(recommendationClient, handler.recommendationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)

	resp, err := recommendationClient.GetRecommendedAccommodations(context.TODO(), &gateway.RecommendationRequest{})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}

	accommodations, err := accommodationClient.PopulateRecommended(context.TODO(), &gateway.PopulateRecommendedRequest{
		Ids: resp.Accommodations,
	})

	if err != nil {
		shared.BadRequest(w, "Something wrong with capturing recommended accommodations!")
		return
	}

	shared.Ok(&w, accommodations.Accommodations)
}
