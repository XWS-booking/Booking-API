package model

import "sort"

type PricingType int32

const (
	PER_GUEST              PricingType = 0
	PER_ACCOMMODATION_UNIT PricingType = 1
)

type Pricing struct {
	Interval    TimeInterval `bson:"interval" json:"interval"`
	Price       float32      `bson:"price" json:"price"`
	PricingType PricingType  `bson:"pricing_type" json:"pricingType"`
}

func (pricing *Pricing) CalculatePrice(interval TimeInterval, guests int32) float32 {
	total := float32(interval.GetDaysDifference()) * pricing.Price
	if pricing.PricingType == PER_GUEST {
		return total * float32(guests)
	}

	return total
}

func AppendPricingIntervals(startingPricing Pricing, intervals []Pricing) (TimeInterval, []Pricing) {
	intervals = SortPricingIntervals(intervals)
	startingInterval := startingPricing.Interval
	var appendedIntervals = []Pricing{startingPricing}

	for _, pricing := range intervals {
		startingInterval.TryAppendInterval(pricing.Interval)
		appendedIntervals = append(appendedIntervals, pricing)
	}
	return startingInterval, appendedIntervals
}

func SortPricingIntervals(intervals []Pricing) []Pricing {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Interval.From.Before(intervals[j].Interval.From)
	})
	return intervals
}
