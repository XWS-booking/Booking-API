package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
	"reservation_service/shared"
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
