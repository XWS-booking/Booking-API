package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
)

type IReservationRepository interface {
	Create(reservation Reservation) (primitive.ObjectID, error)
}
