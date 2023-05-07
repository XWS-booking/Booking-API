package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookingPriceParams struct {
	Interval       TimeInterval       `json:"interval"`
	Guests         int32              `json:"guests"`
	AccomodationId primitive.ObjectID `json:"accomodationId"`
}
