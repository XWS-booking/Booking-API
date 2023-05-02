package model

type Accommodation struct {
	Id             string
	Name           string
	Street         string
	StreetNumber   string
	City           string
	ZipCode        string
	Country        string
	Wifi           bool
	Kitchen        bool
	AirConditioner bool
	FreeParking    bool
	MinGuests      int32
	MaxGuests      int32
	PictureUrls    []string
	OwnerId        string
}

type AccommodationPage struct {
	Data       []Accommodation
	TotalCount int
}
