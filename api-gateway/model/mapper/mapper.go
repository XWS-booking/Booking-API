package mapper

import (
	"gateway/model"
	"gateway/proto/gateway"
)

func UserFromFindUserByIdResponse(resp *gateway.FindUserByIdResponse) model.User {
	return model.User{
		Id:           resp.Id,
		Name:         resp.Name,
		Surname:      resp.Surname,
		Email:        resp.Email,
		Role:         resp.Role,
		Street:       resp.Street,
		StreetNumber: resp.StreetNumber,
		City:         resp.City,
		ZipCode:      resp.Zipcode,
		Country:      resp.Country,
	}
}

func AccommodationFromAccomodationResponse(resp *gateway.AccomodationResponse, owner model.User) model.Accommodation {
	return model.Accommodation{
		Id:             resp.Id,
		Name:           resp.Name,
		Street:         resp.Street,
		StreetNumber:   resp.StreetNumber,
		City:           resp.City,
		ZipCode:        resp.ZipCode,
		Country:        resp.Country,
		Kitchen:        resp.Kitchen,
		Wifi:           resp.Wifi,
		FreeParking:    resp.FreeParking,
		AirConditioner: resp.AirConditioner,
		MinGuests:      resp.MinGuests,
		MaxGuests:      resp.MaxGuests,
		PictureUrls:    resp.Pictures,
		Owner:          owner,
	}
}
