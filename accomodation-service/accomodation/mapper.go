package accomodation

import (
	"accomodation_service/accomodation/dtos"
	"accomodation_service/accomodation/model"
	accomodationGrpc "accomodation_service/proto/accomodation"
	proto "accomodation_service/proto/accomodation"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
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
		Pricing:         mapAccomodationPricing(accomodation.Pricing),
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
		Pricing:         accomodationDto.Pricing,
	}
}

func PricingRequestToPricing(pricingReq proto.Pricing) *model.Pricing {
	return &model.Pricing{
		Price:       pricingReq.Price,
		PricingType: model.PricingType(pricingReq.PricingType),
		Interval: model.TimeInterval{
			From: time.Unix(pricingReq.From.Seconds, int64(pricingReq.From.Nanos)).UTC(),
			To:   time.Unix(pricingReq.To.Seconds, int64(pricingReq.To.Nanos)).UTC(),
		},
	}
}

func mapAccomodationPricing(pricing []model.Pricing) []*accomodationGrpc.Pricing {
	result := make([]*accomodationGrpc.Pricing, 0)
	for _, pric := range pricing {
		result = append(result, &accomodationGrpc.Pricing{
			Price:       pric.Price,
			PricingType: int32(pric.PricingType),
			From:        timestamppb.New(pric.Interval.From),
			To:          timestamppb.New(pric.Interval.To),
		})
	}
	return result
}
