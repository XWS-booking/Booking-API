package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
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
