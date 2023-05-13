package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	"gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	reqCtx "github.com/gorilla/context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strings"
)

type UpdatePricingDto struct {
	Pricing []model.Pricing `bson:"pricing"`
}

type UpdatePricingHandler struct {
	authClientAddress         string
	accomodationClientAddress string
	reservationClientAddress  string
}

func NewUpdatePricingHandler(authClientAddress string, accomodationClientAddress string, reservationClientAddress string) Handler {
	return &UpdatePricingHandler{
		authClientAddress:         authClientAddress,
		accomodationClientAddress: accomodationClientAddress,
		reservationClientAddress:  reservationClientAddress,
	}
}

func (handler *UpdatePricingHandler) Init(mux *runtime.ServeMux) {
	handlerMethod := TokenValidationMiddleware(
		RolesMiddleware(
			[]model.UserRole{model.HOST},
			UserMiddleware(handler.UpdatePricing),
		),
	)
	err := mux.HandlePath("PATCH", "/api/accommodation/{id}", handlerMethod)
	if err != nil {
		panic(err)
	}
}

func (handler *UpdatePricingHandler) UpdatePricing(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {

	var dto UpdatePricingDto
	_ = DecodeBody(r, &dto)
	id := pathParams["id"]
	userId := reqCtx.Get(r, "id").(string)

	//Validate overlapping with reservations
	accomodationClient := services.NewAccommodationClient(handler.accomodationClientAddress)
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)

	pricingsForUpdate := make([]model.Pricing, 0)
	for _, pricing := range dto.Pricing {
		result, _ := reservationClient.IsAccommodationAvailable(context.TODO(), &gateway.IsAccommodationAvailableRequest{
			AccommodationId: id,
			StartDate:       timestamppb.New(pricing.From),
			EndDate:         timestamppb.New(pricing.To),
		})
		if result.Available {
			pricingsForUpdate = append(pricingsForUpdate, pricing)
		}
	}

	if len(pricingsForUpdate) == 0 {
		shared.BadRequest(w, "All intervals have at least one reservation!")
		return
	}
	//Try update
	var pricingMapped []*gateway.Pricing
	for _, p := range pricingsForUpdate {
		mapped := &gateway.Pricing{
			PricingType: p.PricingType,
			Price:       p.Price,
			To:          timestamppb.New(p.To),
			From:        timestamppb.New(p.From),
		}
		pricingMapped = append(pricingMapped, mapped)
	}

	_, err := accomodationClient.UpdatePricing(context.TODO(), &gateway.UpdatePricingRequest{
		Pricing: pricingMapped,
		Id:      id,
		UserId:  userId,
	})

	fmt.Println(err)
	if err != nil {
		fullMessage := strings.Split(err.Error(), "code = ")[1]
		message := strings.Split(fullMessage, " = ")[1]
		shared.BadRequest(w, message)
		return
	}
}
