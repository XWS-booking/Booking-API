package accomodation

import "accomodation_service/accomodation/model"

type IAccomodationRepository interface {
	FindAll(city string, guests int32) ([]model.Accomodation, error)
}
