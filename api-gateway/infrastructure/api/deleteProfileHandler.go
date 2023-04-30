package api

import (
	"context"
	"encoding/json"
	"gateway/infrastructure/services"
	"gateway/proto/gateway"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type DeleteProfileHandler struct {
	authClientAddress          string
	accommodationClientAddress string
	reservationClientAddress   string
}

func NewDeleteProfileHandler(authClientAddress, accommodationClientAddress, reservationClientAddress string) Handler {
	return &DeleteProfileHandler{
		authClientAddress:          authClientAddress,
		accommodationClientAddress: accommodationClientAddress,
		reservationClientAddress:   reservationClientAddress,
	}
}

func (handler *DeleteProfileHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("DELETE", "/api/deleteProfile/{id}/{role}", handler.Delete)
	if err != nil {
		panic(err)
	}
}

func (handler *DeleteProfileHandler) Delete(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := pathParams["id"]
	role := pathParams["role"]
	var canDelete bool
	var err error
	if role == "GUEST" {
		canDelete, err = handler.CanDeleteGuestProfile(id)
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var message string
	if canDelete {
		_, err := handler.DeleteProfile(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		message = "Profile deleted!"
	} else {
		message = "Can't delete profile because there are active reservations!"
	}

	response, err := json.Marshal(message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(response)

}

func (handler *DeleteProfileHandler) CanDeleteGuestProfile(id string) (bool, error) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	activeReservations, err := reservationClient.CheckActiveReservationsForGuest(context.TODO(), &gateway.CheckActiveReservationsForGuestRequest{GuestId: id})
	return !activeReservations.ActiveReservations, err
}

func (handler *DeleteProfileHandler) DeleteProfile(id string) (bool, error) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	deleted, err := authClient.DeleteProfile(context.TODO(), &gateway.DeleteProfileRequest{Id: id})
	return deleted.Deleted, err
}
