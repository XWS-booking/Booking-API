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
	DeleteReservationsByBuyerId(id primitive.ObjectID) error
}
