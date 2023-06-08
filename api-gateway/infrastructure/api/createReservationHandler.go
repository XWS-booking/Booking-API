package api

import (
	"context"
	"fmt"
	"gateway/infrastructure/services"
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
	notificationClientAddress  string
}

func NewCreateReservationHandler(reservationClientAddress, authClientAddress, accommodationClientAddress, notificationClientAddress string) Handler {
	return &CreateReservationHandler{
		reservationClientAddress:   reservationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		authClientAddress:          authClientAddress,
		notificationClientAddress:  notificationClientAddress,
	}
}

func (handler *CreateReservationHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/reservation", handler.Create)
	if err != nil {
		panic(err)
	}
}

func (handler *CreateReservationHandler) Create(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)

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
		http.Error(w, "Failed booking reservation!", http.StatusConflict)
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
			http.Error(w, "Failed to auto confirm!", http.StatusConflict)
		}
	}
	_, err = notificationClient.SendNotification(context.TODO(), &gateway.NotificationRequest{Recipient: accommodation.OwnerId, Message: "You have a new reservation in accommodation '" + accommodation.Name + "'"})
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
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
