package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
	"time"
)

type IReservationRepository interface {
	Create(reservation Reservation) (primitive.ObjectID, error)
	Delete(id primitive.ObjectID) (*primitive.ObjectID, error)
	FindById(id primitive.ObjectID) (Reservation, error)
	FindAllReservedAccommodations(startDate time.Time, endDate time.Time) ([]string, error)
	CheckActiveReservationsForGuest(id primitive.ObjectID) (bool, error)
	CheckActiveReservationsForAccommodation(id primitive.ObjectID) (bool, error)
	DeleteByBuyerId(id primitive.ObjectID) error
	DeleteByAccommodationId(id primitive.ObjectID) error
	UpdateReservation(reservation Reservation) error
	IsAccommodationAvailable(id primitive.ObjectID, startDate time.Time, endDate time.Time) (bool, error)
	FindAllByBuyerId(id primitive.ObjectID) ([]Reservation, error)
	FindNumberOfBuyersCancellations(id primitive.ObjectID) (int, error)
	FindAllByAccommodationId(id primitive.ObjectID) ([]Reservation, error)
	FindAllPendingByAccommodationId(id primitive.ObjectID) ([]Reservation, error)
	CheckIfGuestHasReservationInAccommodation(guestId primitive.ObjectID, accommodationId primitive.ObjectID) (bool, error)
}
