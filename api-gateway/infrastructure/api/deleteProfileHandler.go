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
	err := mux.HandlePath("DELETE", "/api/auth/user", TokenValidationMiddleware(RolesMiddleware([]UserRole{0, 1}, UserMiddleware(handler.Delete))))
	if err != nil {
		panic(err)
	}
}

func (handler *DeleteProfileHandler) Delete(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	token := r.Header["Authorization"][0]

	authClient := services.NewAuthClient(handler.authClientAddress)
	user, e := authClient.GetUser(context.TODO(), &gateway.GetUserRequest{Token: token})
	if e != nil {
		shared.NotFound(w, "User not found")
	}

	var canDelete bool
	var err error
	if user.Role == "0" {
		canDelete, err = handler.CanDeleteGuestProfile(user.Id)
	}
	if user.Role == "1" {
		user, err := authClient.ProfileDeletion(context.TODO(), &gateway.ProfileDeletionRequest{Id: user.Id})
		if err != nil {
			shared.BadRequest(w, err.Error())
			return
		}
		fmt.Println("konacni user", user)
		shared.Ok(&w, user)
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if canDelete {
		deleted, err := handler.DeleteProfile(user.Id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		shared.Ok(&w, deleted)
	} else {
		http.Error(w, "Cannot delete profile due to active reservations!", http.StatusBadRequest)
		return
	}
}

func (handler *DeleteProfileHandler) CanDeleteGuestProfile(id string) (bool, error) {
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	activeReservations, err := reservationClient.CheckActiveReservationsForGuest(context.TODO(), &gateway.CheckActiveReservationsForGuestRequest{GuestId: id})
	return !activeReservations.ActiveReservations, err
}

func (handler *DeleteProfileHandler) CanDeleteHostProfile(id string) (bool, error) {
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	accommodations, err := accommodationClient.FindAllAccommodationIdsByOwnerId(context.TODO(), &gateway.FindAllAccommodationIdsByOwnerIdRequest{OwnerId: id})
	if err != nil {
		return false, err
	}
	activeReservations, err := reservationClient.CheckActiveReservationsForAccommodations(context.TODO(), &gateway.CheckActiveReservationsForAccommodationsRequest{Ids: accommodations.Ids})
	if !activeReservations.ActiveReservations {
		_, err := accommodationClient.DeleteByOwnerId(context.TODO(), &gateway.DeleteByOwnerIdRequest{OwnerId: id})
		if err != nil {
			return false, err
		}
	}
	return !activeReservations.ActiveReservations, err
}

func (handler *DeleteProfileHandler) DeleteProfile(id string) (bool, error) {
	authClient := services.NewAuthClient(handler.authClientAddress)
	deleted, err := authClient.DeleteProfile(context.TODO(), &gateway.DeleteProfileRequest{Id: id})
	return deleted.Deleted, err
}
