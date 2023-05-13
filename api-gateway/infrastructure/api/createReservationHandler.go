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
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/protobuf/types/known/timestamppb"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type CreateReservationDto struct {
	AccommodationId primitive.ObjectID `json:"accommodationId"`
	BuyerId         primitive.ObjectID `json:"buyerId"`
	Guests          string             `json:"guests"`
	StartDate       time.Time          `json:"startDate"`
	EndDate         time.Time          `json:"endDate"`
}

type CreateReservationHandler struct {
	reservationClientAddress   string
	accommodationClientAddress string
	authClientAddress          string
}

func NewCreateReservationHandler(reservationClientAddress, authClientAddress, accommodationClientAddress string) Handler {
	return &CreateReservationHandler{
		reservationClientAddress:   reservationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		authClientAddress:          authClientAddress,
	}
}

func (handler *CreateReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/reservation", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.Create))))
	if err != nil {
		panic(err)
	}
}

func (handler *CreateReservationHandler) Create(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)

	var body CreateReservationDto
	err := DecodeBody(r, &body)
	if err != nil {
		http.Error(w, fmt.Sprintf("failed to parse form: %s", err.Error()), http.StatusBadRequest)
		return
	}
	guests, _ := strconv.ParseInt(body.Guests, 10, 32)
	pricingDto := gateway.GetBookingPriceRequest{
		From:           timestamppb.New(body.StartDate),
		To:             timestamppb.New(body.EndDate),
		Guests:         int32(guests),
		AccomodationId: body.AccommodationId.Hex(),
	}
	price, err := accommodationClient.GetBookingPrice(context.TODO(), &pricingDto)
	if err != nil {
		http.Error(w, "No matching interval for reservation!", http.StatusConflict)
	}

	res, err := reservationClient.Create(context.TODO(), &gateway.CreateReservationRequest{
		AccommodationId: body.AccommodationId.Hex(),
		StartDate:       timestamppb.New(body.StartDate),
		EndDate:         timestamppb.New(body.EndDate),
		Guests:          int32(guests),
		BuyerId:         body.BuyerId.Hex(),
		Price:           price.Price,
	})
	if err != nil {
		panic(err)
	}

	accommodation, err := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: body.AccommodationId.Hex()})
	if err != nil {
		panic(err)
	}
	if accommodation.AutoReservation {
		_, err := reservationClient.Confirm(context.TODO(), &gateway.ReservationId{
			Id: ConvertStringToValidObjectID(res.Id),
		})
		if err != nil {
			panic(err)
		}
	}
	shared.Ok(&w, res)
}

func ConvertStringToValidObjectID(objectIDString string) string {
	start := strings.Index(objectIDString, "\"")
	end := strings.LastIndex(objectIDString, "\"")
	if start == -1 || end == -1 || end <= start {
		fmt.Println("Invalid input string")
		return ""
	}
	return objectIDString[start+1 : end]
}
