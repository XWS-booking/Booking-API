package accomodation

import (
	"accomodation_service/accomodation/model"
)

type AccomodationService struct {
	AccomodationRepository IAccomodationRepository
}

func (accomodationService *AccomodationService) FindAll(city string, guests int32) []model.Accomodation {
	accomodations, err := accomodationService.AccomodationRepository.FindAll(city, guests)
	if err != nil {
		return accomodations
	}
	return accomodations
}
