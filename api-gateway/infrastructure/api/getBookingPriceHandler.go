package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	"gateway/model"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"time"
)

type BookingPriceDto struct {
	Guests int32     `json:"guests"`
	From   time.Time `json:"from"`
	To     time.Time `json:"to"`
}

type GetBookingPriceHandler struct {
	AccommodationClientAddress string
}

func NewGetBookingPriceHandler(accomodationClientAddress string) Handler {
	return &GetBookingPriceHandler{
		AccommodationClientAddress: accomodationClientAddress,
	}
}

func (handler *GetBookingPriceHandler) Init(mux *runtime.ServeMux) {
	handlerMethod := TokenValidationMiddleware(
		RolesMiddleware(
			[]model.UserRole{model.HOST, model.GUEST},
			UserMiddleware(handler.GetBookingPrice),
		),
	)
	err := mux.HandlePath("POST", "/api/accommodation/{id}/price", handlerMethod)
	if err != nil {
		panic(err)
	}
}

func (handler *GetBookingPriceHandler) GetBookingPrice(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	var dto BookingPriceDto
	err := DecodeBody(r, &dto)
	accomodationId := pathParams["id"]
	fmt.Println(err)
	if err != nil {
		shared.BadRequest(w, "Body parameters are wrong")
		return
	}

	accomodationClient := services.NewAccommodationClient(handler.AccommodationClientAddress)
	resp, err := accomodationClient.GetBookingPrice(context.TODO(), &gateway.GetBookingPriceRequest{
		From:           timestamppb.New(dto.From),
		To:             timestamppb.New(dto.To),
		Guests:         dto.Guests,
		AccomodationId: accomodationId,
	})

	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, resp)
}
