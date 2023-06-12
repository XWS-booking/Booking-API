package rate_host

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

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
	Pricing         []Pricing          `bson:"pricing" json:"pricing"`
}

type Pricing struct {
	Uuid        string       `bson:"uuid" json:"uuid"'`
	Interval    TimeInterval `bson:"interval" json:"interval"`
	Price       float32      `bson:"price" json:"price"`
	PricingType PricingType  `bson:"pricing_type" json:"pricingType"`
}

type TimeInterval struct {
	From time.Time `bson:"from" json:"from"`
	To   time.Time `bson:"to" json:"to"`
}

type PricingType int32

type DeleteHostCommandType int8

const (
	FetchHostAccommodationCommand DeleteHostCommandType = iota
	ValidateReservationsCommand
	DeleteAccommodationsCommand
	RemoveHostCommand
	RollbackDeleteAccommodationsCommand
	CancelDeleteHostCommand
	UnknownCommand
)

type DeleteHostCommand struct {
	Host           string
	Accommodations []Accomodation
	Type           DeleteHostCommandType
}

type DeleteHostReplyType int8

const (
	FetchHostAccommodationSuccessReply DeleteHostReplyType = iota
	ValidationSuccessReply
	AccommodationsDeletedReply
	HostDeletedReply
	FetchHostAccommodationFailedReply
	AccommodationDeletionFailedReply
	HostDeletionFailedReply
	RollbackDeleteAccommodationsReply
	CancelDeleteHostReply
	ValidationFailedReply
	UnknownReply
)

type DeleteHostReply struct {
	Host           string
	Accommodations []Accomodation
	Type           DeleteHostReplyType
}
