package model

import (
	"accomodation_service/shared"
	"fmt"
	"github.com/google/uuid"
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

	for i := 0; i <= pricingCount-1; i++ {
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
		fmt.Println(pricing.Interval, interval)
		fmt.Println((interval.From.After(pricing.Interval.From) || interval.From.Equal(pricing.Interval.From)))
		fmt.Println((interval.From.Before(pricing.Interval.To) || interval.From.Equal(pricing.Interval.To)))
		isInside := (interval.From.After(pricing.Interval.From) || interval.From.Equal(pricing.Interval.From)) &&
			(interval.From.Before(pricing.Interval.To) || interval.From.Equal(pricing.Interval.To))
		if isInside {
			return &pricing
		}
	}
	return nil
}

func (accomodation *Accomodation) GeneratePricingUuids() {
	newPricing := make([]Pricing, 0)
	for _, pricing := range accomodation.Pricing {
		if pricing.Uuid == "" {
			id, _ := uuid.NewUUID()
			pricing.Uuid = id.String()
		}
		newPricing = append(newPricing, pricing)
	}
	accomodation.Pricing = newPricing
}

func (accomodation *Accomodation) UpdatePricing(pricing []Pricing) {
	addedPricing := filterAddedPricing(pricing)
	editedPricing := filterEditedPricing(pricing)
	newPricing := make([]Pricing, 0)
	for _, oldPrice := range accomodation.Pricing {
		for i, newPrice := range editedPricing {
			if oldPrice.Uuid == newPrice.Uuid {
				newPricing = append(newPricing, newPrice)
				break
			}
			if i == (len(editedPricing) - 1) {
				newPricing = append(newPricing, oldPrice)
			}
		}
	}
	for _, p := range addedPricing {
		newPricing = append(newPricing, p)
	}
	accomodation.Pricing = newPricing
}

func filterAddedPricing(pricing []Pricing) []Pricing {
	filtered := make([]Pricing, 0)
	for _, price := range pricing {
		if price.Uuid == "" {
			filtered = append(filtered, price)
		}
	}
	return filtered
}
func filterEditedPricing(pricing []Pricing) []Pricing {
	filtered := make([]Pricing, 0)
	for _, price := range pricing {
		if price.Uuid != "" {
			filtered = append(filtered, price)
		}
	}
	return filtered
}
