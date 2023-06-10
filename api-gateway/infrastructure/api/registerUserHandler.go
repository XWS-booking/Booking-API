package api

import (
	"context"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"gateway/shared"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type RegisterUserHandler struct {
	authClientAddress         string
	notificationClientAddress string
}

func NewRegisterUserHandler(authClientAddress, notificationClientAddress string) Handler {
	return &RegisterUserHandler{
		authClientAddress:         authClientAddress,
		notificationClientAddress: notificationClientAddress,
	}
}

func (handler *RegisterUserHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("POST", "/api/auth/register", handler.Register)
	if err != nil {
		panic(err)
	}
}

func (handler *RegisterUserHandler) Register(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	notificationClient := services.NewNotificationClient(handler.notificationClientAddress)
	var body *gateway.RegistrationRequest
	err := DecodeBody(r, &body)
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	registrationRes, err := authClient.Register(context.TODO(), body)
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	request := &gateway.CreateNotificationPreferencesRequest{UserId: registrationRes.Id,
		GuestCreatedReservationRequest:     registrationRes.Role == "1",
		GuestCanceledReservation:           registrationRes.Role == "1",
		GuestRatedHost:                     registrationRes.Role == "1",
		GuestRatedAccommodation:            registrationRes.Role == "1",
		DistinguishedHost:                  registrationRes.Role == "1",
		HostConfirmedOrRejectedReservation: registrationRes.Role == "0",
	}
	_, err = notificationClient.CreateNotificationPreferences(context.TODO(), request)
	if err != nil {
		shared.BadRequest(w, err.Error())
		return
	}
	shared.Ok(&w, registrationRes)
}
