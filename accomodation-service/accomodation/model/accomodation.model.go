package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Accomodation struct {
	Id              primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name            string             `bson:"name" json:"name"`
	Street          string             `bson:"street" json:"street"`
	StreetNumber    string             `bson:"street_number" json:"streetNumber"`
	City            string             `bson:"city" json:"city"`
	ZipCode         string             `bson:"zip_code" json:"zipCode"`
	Country         string             `bson:"country" json:"country"`
	Wifi            bool               `bson:"wifi" json:"wifi"`
	Kitchen         bool               `bson:"kitchen" json:"kitchen"`
	AirConditioner  bool               `bson:"air_conditioner" json:"airConditioner"`
	FreeParking     bool               `bson:"free_parking" json:"freeParking"`
	AutoReservation bool               `bson:"auto_reservation" json:"autoReservation"`
	MinGuests       int32              `bson:"min_guests" json:"minGuests"`
	MaxGuests       int32              `bson:"max_guests" json:"maxGuests"`
	PictureUrls     []string           `bson:"picture_urls" json:"pictureUrls"`
	OwnerId         primitive.ObjectID `bson:"owner_id" json:"ownerId"`
}
