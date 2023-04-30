package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
	"reservation_service/shared"
	"time"
)

type ReservationService struct {
	ReservationRepository IReservationRepository
}

func (reservationService *ReservationService) Create(reservation Reservation) primitive.ObjectID {
	reservation.Status = 0
	created, error := reservationService.ReservationRepository.Create(reservation)
	if error != nil {
		return created
	}
	return created
}

func (reservationService *ReservationService) Delete(id primitive.ObjectID) *shared.Error {
	_, e := reservationService.ReservationRepository.FindById(id)
	if e != nil {
		return shared.ReservationNotFound()
	}
	error := reservationService.ReservationRepository.Delete(id)
	if error != nil {
		return shared.ReservationNotDeleted()
	}
	return nil
}

func (reservationService ReservationService) FindAllReservedAccommodations(startDate time.Time, endDate time.Time) ([]string, *shared.Error) {
	accommodationIds, err := reservationService.ReservationRepository.FindAllReservedAccommodations(startDate, endDate)
	if err != nil {
		return accommodationIds, shared.ReservationsNotFound()
	}
	return accommodationIds, nil
}

func (reservationService *ReservationService) CheckActiveReservationsForGuest(id primitive.ObjectID) (bool, *shared.Error) {
	activeReservations, err := reservationService.ReservationRepository.CheckActiveReservationsForGuest(id)
	if err != nil {
		return activeReservations, shared.CheckActiveReservationsError()
	}
	err = reservationService.ReservationRepository.DeleteReservationsByBuyerId(id)
	if err != nil {
		return activeReservations, shared.ReservationNotDeleted()
	}
	return activeReservations, nil
}
