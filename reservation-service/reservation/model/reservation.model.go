package model

import (
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/proto/reservation"
	. "reservation_service/shared"
	"time"
)

type Status = int

const (
	Pending Status = iota
	Confirmed
	Rejected
	Canceled
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodationId"`
	BuyerId         primitive.ObjectID `bson:"buyer_id" json:"buyerId"`
	StartDate       time.Time          `bson:"start_date" json:"startDate"`
	EndDate         time.Time          `bson:"end_date" json:"endDate"`
	Guests          int32              `bson:"guests" json:"guests"`
	Price           float32            `bson:"price" json:"price"`
	Status          Status             `bson:"status" json:"status"`
}

func NewReservation(req *CreateReservationRequest) Reservation {
	accommodationId, _ := primitive.ObjectIDFromHex(req.AccommodationId)
	buyerId, _ := primitive.ObjectIDFromHex(req.BuyerId)
	startDate, _ := ptypes.Timestamp(req.StartDate)
	endDate, _ := ptypes.Timestamp(req.EndDate)
	return Reservation{
		AccommodationId: accommodationId,
		StartDate:       startDate,
		EndDate:         endDate,
		Guests:          req.Guests,
		BuyerId:         buyerId,
		Price:           req.Price,
	}
}

func (reservation *Reservation) Cancel() *Error {
	if reservation.Status != 1 {
		return ReservationCancelationFailed()
	}
	if time.Now().Add(time.Hour * 24).After(reservation.StartDate) {
		return ReservationCancelationTooLate()
	}
	reservation.Status = Canceled
	return nil
}

func (reservation *Reservation) Confirm() *Error {
	if reservation.Status != 0 {
		return ReservationConfirmationFailed()
	}
	reservation.Status = Confirmed
	return nil
}

func (reservation *Reservation) Reject() *Error {
	if reservation.Status != 0 {
		return ReservationRejectionFailed()
	}
	reservation.Status = Rejected
	return nil
}
