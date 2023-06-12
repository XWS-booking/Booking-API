package handlers

import (
	"accomodation_service/accomodation/model"
	events "accomodation_service/common/rate_host"
)

func mapAccommodations(accommodations []model.Accomodation) []events.Accomodation {
	mappedAccommodations := make([]events.Accomodation, 0)
	for _, accommodation := range accommodations {
		mappedAccommodations = append(mappedAccommodations, events.Accomodation{
			Id:              accommodation.Id,
			Name:            accommodation.Name,
			Street:          accommodation.Street,
			StreetNumber:    accommodation.StreetNumber,
			City:            accommodation.City,
			ZipCode:         accommodation.ZipCode,
			Country:         accommodation.Country,
			Wifi:            accommodation.Wifi,
			Kitchen:         accommodation.Kitchen,
			AirConditioner:  accommodation.AirConditioner,
			FreeParking:     accommodation.FreeParking,
			AutoReservation: accommodation.AutoReservation,
			MinGuests:       accommodation.MinGuests,
			MaxGuests:       accommodation.MaxGuests,
			PictureUrls:     accommodation.PictureUrls,
			OwnerId:         accommodation.OwnerId,
			Pricing:         mapPricing(accommodation.Pricing),
		})
	}

	return mappedAccommodations
}

func mapPricing(pricings []model.Pricing) []events.Pricing {
	mappedPricings := make([]events.Pricing, 0)
	for _, pricing := range pricings {
		mappedPricings = append(mappedPricings, events.Pricing{
			Uuid: pricing.Uuid,
			Interval: events.TimeInterval{
				From: pricing.Interval.From,
				To:   pricing.Interval.To,
			},

			Price:       pricing.Price,
			PricingType: events.PricingType(pricing.PricingType),
		})
	}
	return mappedPricings
}

func mapPricingEventToModel(pricings []events.Pricing) []model.Pricing {
	mappedPricings := make([]model.Pricing, 0)
	for _, pricing := range pricings {
		mappedPricings = append(mappedPricings, model.Pricing{
			Uuid: pricing.Uuid,
			Interval: model.TimeInterval{
				From: pricing.Interval.From,
				To:   pricing.Interval.To,
			},

			Price:       pricing.Price,
			PricingType: model.PricingType(pricing.PricingType),
		})
	}
	return mappedPricings
}
