package dtos

import (
	"accomodation_service/accomodation/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccomodationDto struct {
	Id             string             `json:"id"`
	Name           string             `json:"name"`
	Street         string             `json:"street"`
	StreetNumber   string             `json:"streetNumber"`
	City           string             `json:"city"`
	ZipCode        string             `json:"zipCode"`
	Country        string             `json:"country"`
	Wifi           bool               `json:"wifi"`
	Kitchen        bool               `json:"kitchen"`
	AirConditioner bool               `json:"airConditioner"`
	FreeParking    bool               `json:"freeParking"`
	MinGuests      int32              `json:"minGuests"`
	MaxGuests      int32              `json:"maxGuests"`
	Pictures       []byte             `json:"pictures"`
	OwnerId        primitive.ObjectID `json:"ownerId"`
}

func NewAccomodationDto(accomodation model.Accomodation) *AccomodationDto {
	return &AccomodationDto{
		Id:             accomodation.Id.Hex(),
		Name:           accomodation.Name,
		Street:         accomodation.Street,
		StreetNumber:   accomodation.StreetNumber,
		City:           accomodation.City,
		ZipCode:        accomodation.ZipCode,
		Country:        accomodation.Country,
		Wifi:           accomodation.Wifi,
		Kitchen:        accomodation.Kitchen,
		AirConditioner: accomodation.AirConditioner,
		FreeParking:    accomodation.FreeParking,
		MinGuests:      accomodation.MinGuests,
		MaxGuests:      accomodation.MaxGuests,
		OwnerId:        accomodation.OwnerId,
	}

}
