package model

import (
	"accomodation_service/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type PricingCalculationParams struct {
	fromDate time.Time `json:"fromDate"`
	toDate   time.Time `json:"toDate"`
	guests   int32     `json:"guests"`
}

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

func (accomodation *Accomodation) CalculateBookingPrice(interval TimeInterval, guests int32) (float32, *shared.Error) {
	startingPricing := accomodation.FindStartingPricingInterval(interval)

	if startingPricing == nil {
		return 0, shared.NoMatchingPricingInterval()
	}

	resultInterval, appendedPricing := AppendPricingIntervals(*startingPricing, accomodation.Pricing)
	if !resultInterval.HasTimeIntervalInside(interval) {
		return 0, shared.NoMatchingPricingInterval()
	}

	var totalPrice float32 = 0
	for _, pricing := range appendedPricing {
		totalPrice += pricing.CalculatePrice(interval, guests)
	}
	return totalPrice, nil
}

func (accomodation *Accomodation) ValidatePricing() *shared.Error {
	pricingCount := len(accomodation.Pricing)
	if pricingCount == 0 {
		return shared.NoPricingIntervalsFound()
	}
	sorted := SortPricingIntervals(accomodation.Pricing)

	for i := 1; i <= pricingCount-1; i++ {
		if i == (pricingCount - 1) {
			break
		}
		if sorted[i].Interval.IsOverlapping(sorted[i+1].Interval) {
			return shared.PricingIntervalsOverlapping()
		}
	}
	return nil
}

func (accomodation *Accomodation) GetTimeIntervalsFromPricing() []TimeInterval {
	result := make([]TimeInterval, 0)
	for _, pricing := range accomodation.Pricing {
		result = append(result, pricing.Interval)
	}
	return result
}

func (accomodation *Accomodation) FindStartingPricingInterval(interval TimeInterval) *Pricing {
	for _, pricing := range accomodation.Pricing {
		if pricing.Interval.IsOverlapping(interval) {
			return &pricing
		}
	}
	return nil
}
