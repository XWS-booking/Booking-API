package accomodation

import (
	"accomodation_service/accomodation/model"
	accomodationGrpc "accomodation_service/proto/accomodation"
)

func NewAccomodationResponse(accomodation model.Accomodation) *accomodationGrpc.AccomodationResponse {
	return &accomodationGrpc.AccomodationResponse{
		Id:             accomodation.Id.String(),
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
	}
}
