package model

import (
	"time"
)

type Rating struct {
	Id              string `json:"id"`
	AccommodationId string `json:"accommodationId"`
	GuestId         string `json:"guestId"`
	Rating          int32  `json:"rating"`
}

type Reservation struct {
	Id                  string        `json:"id"`
	Accommodation       Accommodation `json:"accommodation"`
	BuyerId             string        `json:"buyerId"`
	StartDate           time.Time     `json:"startDate"`
	EndDate             time.Time     `json:"endDate"`
	Guests              int32         `json:"guests"`
	Status              int32         `json:"status"`
	AccommodationRating Rating        `json:"accommodationRating"`
}
