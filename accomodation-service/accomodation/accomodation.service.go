package accomodation

import (
	"accomodation_service/accomodation/dtos"
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

func (accomodationService *AccomodationService) SearchAndFilter(params dtos.SearchDto) (dtos.SearchResultDto, *shared.Error) {
	accomodations, err := accomodationService.AccomodationRepository.SearchAndFilter(params)
	if err != nil {
		return dtos.SearchResultDto{Data: []model.Accomodation{}, TotalCount: 0}, nil
	}
	totalCount := accomodationService.AccomodationRepository.CountTotalForSearchAndFilter(params)
	return dtos.SearchResultDto{Data: accomodations, TotalCount: totalCount}, nil

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
	accomodation.GeneratePricingUuids()
	err := accomodation.ValidatePricing()
	if err != nil {
		return nil, err
	}

	created, e := accomodationService.AccomodationRepository.Create(accomodation)
	if e != nil {
		return nil, shared.AccomodationNotCreated()
	}
	return created, nil
}

func (accomodationService *AccomodationService) FindById(id primitive.ObjectID) (model.Accomodation, *shared.Error) {
	accommodation, e := accomodationService.AccomodationRepository.FindById(id)
	if e != nil {
		return accommodation, shared.AccommodationsNotFound()
	}
	return accommodation, nil
}

func (accomodationService *AccomodationService) GetBookingPrice(params model.BookingPriceParams) (float32, *shared.Error) {
	accomodation, err := accomodationService.FindById(params.AccomodationId)
	if err != nil {
		return 0, err
	}
	price, err := accomodation.CalculateBookingPrice(params.Interval, params.Guests)
	if err != nil {
		return 0, err
	}
	return price, nil
}

func (accomodationService *AccomodationService) UpdatePricing(acc model.Accomodation) *shared.Error {
	accomodation, err := accomodationService.FindById(acc.Id)
	if err != nil {
		return shared.AccommodationsNotFound()
	}
	if acc.OwnerId != accomodation.OwnerId {
		return shared.NotAccomodationOwner()
	}
	accomodation.UpdatePricing(acc.Pricing)
	accomodation.GeneratePricingUuids()
	err = accomodation.ValidatePricing()
	if err != nil {
		return err
	}
	accomodationService.AccomodationRepository.UpdatePricing(accomodation)
	return nil
}

func (accommodationService *AccomodationService) PopulateRecommended(ids []primitive.ObjectID) ([]model.Accomodation, error) {
	accommodations, err := accommodationService.AccomodationRepository.FindAllByIds(ids)
	return accommodations, err
}
