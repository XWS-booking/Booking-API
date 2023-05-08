package model

type User struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	Role         int32  `json:"role"`
	Street       string `json:"street"`
	StreetNumber string `json:"streetNumber"`
	City         string `json:"city"`
	ZipCode      string `json:"zipCode"`
	Country      string `json:"country"`
}

type UserRole int32

const (
	GUEST             UserRole = 0
	HOST              UserRole = 1
	NOT_AUTHENTICATED UserRole = 2
)
