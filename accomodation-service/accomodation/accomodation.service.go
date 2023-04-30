package accomodation

import (
	"accomodation_service/accomodation/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (accomodationService *AccomodationService) FindAllByOwnerId(id primitive.ObjectID) []model.Accomodation {
	accomodations, err := accomodationService.AccomodationRepository.FindAllByOwnerId(id)
	if err != nil {
		return accomodations
	}
	return accomodations
}

func (accomodationService *AccomodationService) DeleteByOwnerId(id primitive.ObjectID) {
	accomodationService.AccomodationRepository.DeleteByOwnerId(id)
}
