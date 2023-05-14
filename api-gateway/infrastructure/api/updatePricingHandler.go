package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	"gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	"gateway/shared"
	reqCtx "github.com/gorilla/context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strings"
	"time"
)

type UpdatePricingDto struct {
	Pricing []model.Pricing `bson:"pricing"`
}
type TimeInterval struct {
	From time.Time
	To   time.Time
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
	accomodationResp, _ := accomodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{
		Id: id,
	})
	accomodation := mapper.AccommodationFromAccomodationResponseWithoutUser(accomodationResp)

	pricingFiltered := make([]model.Pricing, 0)
	for _, pricing := range dto.Pricing {
		availabilityRequest := &gateway.IsAccommodationAvailableRequest{}
		if pricing.Uuid != "" {
			existingPricing := accomodation.FindPricingByUuid(pricing.Uuid)
			if existingPricing == nil {
				continue
			}
			availabilityRequest = &gateway.IsAccommodationAvailableRequest{
				AccommodationId: id,
				StartDate:       timestamppb.New(existingPricing.From),
				EndDate:         timestamppb.New(existingPricing.To),
			}
			result, _ := reservationClient.IsAccommodationAvailable(context.TODO(), availabilityRequest)
			if !result.Available {
				continue
			}
		}
		availabilityRequest = &gateway.IsAccommodationAvailableRequest{
			AccommodationId: id,
			StartDate:       timestamppb.New(pricing.From),
			EndDate:         timestamppb.New(pricing.To),
		}

		result, _ := reservationClient.IsAccommodationAvailable(context.TODO(), availabilityRequest)
		if result.Available {
			pricingFiltered = append(pricingFiltered, pricing)
		}
	}

	//Try update
	var pricingMapped []*gateway.Pricing
	for _, p := range pricingFiltered {
		mapped := &gateway.Pricing{
			PricingType: p.PricingType,
			Price:       p.Price,
			To:          timestamppb.New(p.To),
			From:        timestamppb.New(p.From),
			Uuid:        p.Uuid,
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

func (timeInterval *TimeInterval) IsOverlapping(interval TimeInterval) bool {
	start := max(interval.From, timeInterval.From)
	end := min(interval.To, timeInterval.To)
	return !(start.After(end) || start.Equal(end))
}

func max(time1 time.Time, time2 time.Time) time.Time {
	if time1.After(time2) {
		return time1
	}
	return time2
}

func min(time1 time.Time, time2 time.Time) time.Time {
	if time1.Before(time2) {
		return time1
	}
	return time2
}

func (timeInterval *TimeInterval) HasTimeIntervalInside(interval TimeInterval) bool {
	return (interval.From.After(timeInterval.From) || interval.From.Equal(timeInterval.From)) &&
		(interval.To.Before(timeInterval.To) || interval.To.Equal(timeInterval.To))
}
