package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
	"time"
)

type IReservationRepository interface {
	Create(reservation Reservation) (primitive.ObjectID, error)
	Delete(id primitive.ObjectID) error
	FindById(id primitive.ObjectID) (Reservation, error)
	FindAllReservedAccommodations(startDate time.Time, endDate time.Time) ([]string, error)
	CheckActiveReservationsForGuest(id primitive.ObjectID) (bool, error)
	CheckActiveReservationsForAccommodation(id primitive.ObjectID) (bool, error)
	DeleteByBuyerId(id primitive.ObjectID) error
	DeleteByAccommodationId(id primitive.ObjectID) error
}
