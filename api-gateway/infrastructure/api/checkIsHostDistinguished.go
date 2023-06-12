package api

import (
	"context"
	"gateway/infrastructure/services"
	. "gateway/middlewares"
	. "gateway/model"
	"gateway/model/mapper"
	"gateway/proto/gateway"
	ctx "github.com/gorilla/context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"net/http"
)

type CheckIsHostDistinguishedHandler struct {
	ratingClientAddress        string
	accommodationClientAddress string
	reservationClientAddress   string
}

func newCheckIsHostDistinguishedHandler(ratingClientAddress string, reservationClientAddress string, accommodationClientAddress string) Handler {
	return &CheckIsHostDistinguishedHandler{
		ratingClientAddress:        ratingClientAddress,
		reservationClientAddress:   reservationClientAddress,
		accommodationClientAddress: accommodationClientAddress,
	}
}

func (handler *CheckIsHostDistinguishedHandler) Init(mux *runtime.ServeMux) {
	err := mux.HandlePath("GET", "/api/check/distinguished-host", TokenValidationMiddleware(RolesMiddleware([]UserRole{0}, UserMiddleware(handler.CheckIsHostDistinguished))))
	if err != nil {
		panic(err)
	}
}

func (handler *CheckIsHostDistinguishedHandler) CheckIsHostDistinguished(w http.ResponseWriter, r *http.Request, pathParams map[string]string) {
	id := ctx.Get(r, "id").(string)
	ratingClient := services.NewRatingClient(handler.ratingClientAddress)
	reservations := handler.FindAllReservationsByOwner(id)
	hostRating, e := ratingClient.GetAverageHostRating(context.TODO(), &gateway.GetAverageHostRatingRequest{HostId: id})
	if e != nil {
		http.Error(w, "Problem with finding host average rating!", http.StatusBadRequest)
	}
	isThereLessThan5PercentCanceledReservations := handler.CheckIsThereLessThan5PercentCanceledReservations(reservations)
	isThereMoreThan50DaysOfReservations := handler.CheckIsThereMoreThan50DaysOfReservations(reservations)
	isThereMoreThan5Reservations := handler.CheckIsThereMoreThan5Reservations(reservations)
	isHostDistinguished := hostRating.Rating > 4.7 &&
		isThereLessThan5PercentCanceledReservations &&
		isThereMoreThan50DaysOfReservations &&
		isThereMoreThan5Reservations
	
}

func (handler *CheckIsHostDistinguishedHandler) CheckIsThereLessThan5PercentCanceledReservations(reservations []ReservationWithCancellation) bool {
	canceledReservations := make([]ReservationWithCancellation, 0)

	for _, reservation := range reservations {
		if reservation.Status == 3 {
			canceledReservations = append(canceledReservations, reservation)
		}
	}
	numberOfCanceledReservations := float64(len(canceledReservations))
	numberOfReservations := float64(len(reservations))
	return numberOfCanceledReservations/numberOfReservations*100 < 5
}

func (handler *CheckIsHostDistinguishedHandler) CheckIsThereMoreThan50DaysOfReservations(reservations []ReservationWithCancellation) bool {
	totalDays := 0

	for _, reservation := range reservations {
		duration := reservation.EndDate.Sub(reservation.StartDate)
		days := int(duration.Hours() / 24)
		totalDays += days
	}
	return totalDays > 50
}

func (handler *CheckIsHostDistinguishedHandler) CheckIsThereMoreThan5Reservations(reservations []ReservationWithCancellation) bool {
	return len(reservations) > 5
}

func (handler *CheckIsHostDistinguishedHandler) FindAllReservationsByOwner(id string) []ReservationWithCancellation {
	accommodationClient := services.NewAccommodationClient(handler.accommodationClientAddress)
	accommodations, e := accommodationClient.FindAllAccommodationIdsByOwnerId(context.TODO(), &gateway.FindAllAccommodationIdsByOwnerIdRequest{OwnerId: id})
	if e != nil {
		return []ReservationWithCancellation{}
	}
	reservationClient := services.NewReservationClient(handler.reservationClientAddress)
	var reservationsWithAccommodation []ReservationWithCancellation
	for _, accommId := range accommodations.Ids {
		accommodation, e := accommodationClient.FindById(context.TODO(), &gateway.FindAccommodationByIdRequest{Id: accommId})
		if e != nil {
			return []ReservationWithCancellation{}
		}
		reservations, e := reservationClient.FindAllByAccommodationId(context.TODO(), &gateway.FindAllReservationsByAccommodationIdRequest{AccommodationId: accommId})
		if e != nil {
			return []ReservationWithCancellation{}
		}
		for _, reservation := range reservations.Reservations {
			numberOfCancellation, e := reservationClient.FindNumberOfBuyersCancellations(context.TODO(), &gateway.NumberOfCancellationRequest{BuyerId: reservation.BuyerId})
			if e != nil {
				return []ReservationWithCancellation{}
			}
			reservationsWithAccommodation = append(reservationsWithAccommodation, ReservationWithCancellation{
				Id:                   reservation.Id,
				Accommodation:        mapper.AccommodationFromAccomodationResponse(accommodation, User{}, 0),
				BuyerId:              reservation.BuyerId,
				StartDate:            reservation.StartDate.AsTime(),
				EndDate:              reservation.EndDate.AsTime(),
				Guests:               reservation.Guests,
				Status:               reservation.Status,
				NumberOfCancellation: numberOfCancellation.CancellationNumber,
			})
		}
	}
	return reservationsWithAccommodation
}
