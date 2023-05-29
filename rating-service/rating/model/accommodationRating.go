package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccommodationRating struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	AccommodationId primitive.ObjectID `bson:"accommodation_id" json:"accommodationId"`
	GuestId         primitive.ObjectID `bson:"guest_id" json:"guestId"`
	Rating          int32              `bson:"rating" json:"rating"`
}
