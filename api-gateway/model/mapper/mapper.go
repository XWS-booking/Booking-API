package mapper

import (
	"gateway/model"
	"gateway/proto/gateway"
	"time"
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
		Pricing:        mapPricingResponseToPricing(resp.Pricing),
	}
}

func mapPricingResponseToPricing(pricing []*gateway.Pricing) []model.Pricing {
	result := make([]model.Pricing, 0)
	for _, prc := range pricing {
		from := time.Unix(prc.From.Seconds, int64(prc.From.Nanos)).UTC()
		to := time.Unix(prc.To.Seconds, int64(prc.To.Nanos)).UTC()
		result = append(result, model.Pricing{
			From:        from,
			To:          to,
			Price:       prc.Price,
			PricingType: prc.PricingType,
		})
	}
	return result
}
