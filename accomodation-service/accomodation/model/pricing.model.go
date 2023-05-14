package model

import (
	"fmt"
	"sort"
)

type PricingType int32

const (
	PER_GUEST              PricingType = 0
	PER_ACCOMMODATION_UNIT PricingType = 1
)

type Pricing struct {
	Uuid        string       `bson:"uuid" json:"uuid"'`
	Interval    TimeInterval `bson:"interval" json:"interval"`
	Price       float32      `bson:"price" json:"price"`
	PricingType PricingType  `bson:"pricing_type" json:"pricingType"`
}

func (pricing *Pricing) CalculatePrice(interval TimeInterval, guests int32) float32 {
	newInterval := TimeInterval{
		From: max(pricing.Interval.From, interval.From),
		To:   min(pricing.Interval.To, interval.To),
	}
	fmt.Println("This is new interval for payment", newInterval)
	fmt.Println("This is date difference", newInterval.GetDaysDifference())
	if newInterval.From.After(newInterval.To) {
		return 0
	}
	total := float32(newInterval.GetDaysDifference()+1) * pricing.Price
	if pricing.PricingType == PER_GUEST {
		return total * float32(guests)
	}
	fmt.Println("This is pricing", pricing.Price)
	fmt.Println("This is total price", total)
	return total
}

func AppendPricingIntervals(startingPricing Pricing, intervals []Pricing) (TimeInterval, []Pricing) {
	intervals = SortPricingIntervals(intervals)
	startingInterval := startingPricing.Interval
	var appendedIntervals = []Pricing{startingPricing}

	for _, pricing := range intervals {
		isAppended := startingInterval.TryAppendInterval(pricing.Interval)
		if isAppended {
			appendedIntervals = append(appendedIntervals, pricing)
		}
	}
	return startingInterval, appendedIntervals
}

func SortPricingIntervals(intervals []Pricing) []Pricing {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i].Interval.From.Before(intervals[j].Interval.From)
	})
	return intervals
}
