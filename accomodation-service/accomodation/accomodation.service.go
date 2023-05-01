package accomodation

import (
	"accomodation_service/accomodation/model"
	shared "accomodation_service/shared"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AccomodationService struct {
	AccomodationRepository IAccomodationRepository
}

func (accomodationService *AccomodationService) FindAll(city string, guests int32) ([]model.Accomodation, *shared.Error) {
	accomodations, err := accomodationService.AccomodationRepository.FindAll(city, guests)
	if err != nil {
		return accomodations, shared.AccommodationsNotFound()
	}
	return accomodations, nil
}

func (accomodationService *AccomodationService) FindAllByOwnerId(id primitive.ObjectID) ([]model.Accomodation, *shared.Error) {
	accomodations, err := accomodationService.AccomodationRepository.FindAllByOwnerId(id)
	if err != nil {
		return accomodations, shared.AccommodationsNotFound()
	}
	return accomodations, nil
}

func (accomodationService *AccomodationService) DeleteByOwnerId(id primitive.ObjectID) *shared.Error {
	err := accomodationService.AccomodationRepository.DeleteByOwnerId(id)
	if err != nil {
		return shared.AccommodationNotDeleted()
	}
	return nil
}

func (accomodationService *AccomodationService) Create(accomodation model.Accomodation) (*model.Accomodation, *shared.Error) {
	created, e := accomodationService.AccomodationRepository.Create(accomodation)
	if e != nil {
		return nil, shared.AccomodationNotCreated()
	}
	return created, nil
}
