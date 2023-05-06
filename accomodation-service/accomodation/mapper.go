package accomodation

import (
	"accomodation_service/accomodation/dtos"
	"accomodation_service/accomodation/model"
	accomodationGrpc "accomodation_service/proto/accomodation"
)

func NewAccomodationResponse(accomodation model.Accomodation) *accomodationGrpc.AccomodationResponse {
	return &accomodationGrpc.AccomodationResponse{
		Id:              accomodation.Id.Hex(),
		Name:            accomodation.Name,
		Street:          accomodation.Street,
		StreetNumber:    accomodation.StreetNumber,
		City:            accomodation.City,
		ZipCode:         accomodation.ZipCode,
		Country:         accomodation.Country,
		Wifi:            accomodation.Wifi,
		Kitchen:         accomodation.Kitchen,
		AirConditioner:  accomodation.AirConditioner,
		AutoReservation: accomodation.AutoReservation,
		FreeParking:     accomodation.FreeParking,
		MinGuests:       accomodation.MinGuests,
		MaxGuests:       accomodation.MaxGuests,
		Pictures:        accomodation.PictureUrls,
		OwnerId:         accomodation.OwnerId.Hex(),
	}
}

func AccomodationDtoToAccomodation(accomodationDto dtos.AccomodationDto) model.Accomodation {
	return model.Accomodation{
		Name:            accomodationDto.Name,
		Street:          accomodationDto.Street,
		StreetNumber:    accomodationDto.StreetNumber,
		City:            accomodationDto.City,
		ZipCode:         accomodationDto.ZipCode,
		Country:         accomodationDto.Country,
		Wifi:            accomodationDto.Wifi,
		Kitchen:         accomodationDto.Kitchen,
		AirConditioner:  accomodationDto.AirConditioner,
		AutoReservation: accomodationDto.AutoReservation,
		FreeParking:     accomodationDto.FreeParking,
		MinGuests:       accomodationDto.MinGuests,
		MaxGuests:       accomodationDto.MaxGuests,
		OwnerId:         accomodationDto.OwnerId,
	}
}
