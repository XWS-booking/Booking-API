package dtos

import "accomodation_service/accomodation/model"

type PriceRange struct {
	From float64
	To   float64
}

type SearchDto struct {
	IncludingIds []string
	City         string
	Guests       int32
	Filters      []string
	Price        PriceRange
	Page         int32
	Limit        int32
}

type SearchResultDto struct {
	Data       []model.Accomodation
	TotalCount int32
}
