package model

import "time"

type ReservationWithCancellation struct {
	Id                   string        `json:"id"`
	Accommodation        Accommodation `json:"accommodation"`
	BuyerId              string        `json:"buyerId"`
	StartDate            time.Time     `json:"startDate"`
	EndDate              time.Time     `json:"endDate"`
	Guests               int32         `json:"guests"`
	Status               int32         `json:"status"`
	NumberOfCancellation int32         `json:"numberOfCancellation"`
}
