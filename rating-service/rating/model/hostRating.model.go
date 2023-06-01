package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type HostRating struct {
	Id      primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	HostId  primitive.ObjectID `bson:"host_id" json:"hostId"`
	GuestId primitive.ObjectID `bson:"guest_id" json:"guestId"`
	Rating  int32              `bson:"rating" json:"rating"`
	Time    time.Time          `bson:"time" json:"time"`
}
