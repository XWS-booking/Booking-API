package dtos

import (
	"accomodation_service/accomodation/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PriceRange struct {
	From float64
	To   float64
}

type SearchDto struct {
	IncludingIds    []primitive.ObjectID
	City            string
	Guests          int32
	Filters         []string
	Price           PriceRange
	Page            int32
	Limit           int32
	FeaturedHostIds []primitive.ObjectID
}

type SearchResultDto struct {
	Data       []model.Accomodation
	TotalCount int32
}
