package model

import "time"

type Rating struct {
	Id              string    `json:"id"`
	Value           int32     `json:"value"`
	CreatedAt       time.Time `json:"createdAt"`
	UserId          string    `json:"userId"`
	AccommodationId string    `json:"accommodationId"`
}
