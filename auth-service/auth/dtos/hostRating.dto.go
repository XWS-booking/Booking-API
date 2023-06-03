package dtos

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HostRatingDto struct {
	Id      primitive.ObjectID `json:"id"`
	GuestId primitive.ObjectID `json:"guestId"`
	HostId  primitive.ObjectID `json:"hostId"`
	Rating  int                `json:"rating"`
	Time    time.Time          `json:"time"`
}
