package model

type Accommodation struct {
	Id              string   `json:"id"`
	Name            string   `json:"name"`
	Street          string   `json:"street"`
	StreetNumber    string   `json:"streetNumber"`
	City            string   `json:"city"`
	ZipCode         string   `json:"zipCode"`
	Country         string   `json:"country"`
	Wifi            bool     `json:"wifi"`
	Kitchen         bool     `json:"kitchen"`
	AirConditioner  bool     `json:"airConditioner"`
	AutoReservation bool     `json:"autoReservation"`
	FreeParking     bool     `json:"freeParking"`
	MinGuests       int32    `json:"minGuests"`
	MaxGuests       int32    `json:"maxGuests"`
	PictureUrls     []string `json:"pictureUrls"`
	Owner           User     `json:"owner"`
}

type AccommodationPage struct {
	Data       []Accommodation `json:"data"`
	TotalCount int             `json:"totalCount"`
}
