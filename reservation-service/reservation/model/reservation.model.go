package model

import (
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/proto/reservation"
	"time"
)

type Status = int

const (
	Pending Status = iota
	Approved
	Rejected
)

type Reservation struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodationId"`
	StartDate       time.Time          `bson:"start_date" json:"startDate"`
	EndDate         time.Time          `bson:"end_date" json:"endDate"`
	guests          int32              `bson:"guests" json:"guests"`
	Status          Status             `bson:"status" json:"status"`
}

func NewReservation(req *CreateReservationRequest) Reservation {
	accommodationId, _ := primitive.ObjectIDFromHex(req.AccommodationId)
	startDate, _ := ptypes.Timestamp(req.StartDate)
	endDate, _ := ptypes.Timestamp(req.EndDate)
	return Reservation{
		AccommodationId: accommodationId,
		StartDate:       startDate,
		EndDate:         endDate,
		guests:          req.Guests,
	}
}
