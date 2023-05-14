package accomodation

import (
	. "accomodation_service/accomodation/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IAccomodationRepository interface {
	FindAll(city string, guests int32) ([]Accomodation, error)
	FindAllByOwnerId(id primitive.ObjectID) ([]Accomodation, error)
	DeleteByOwnerId(id primitive.ObjectID) error
	Create(accomodation Accomodation) (*Accomodation, error)
	FindById(id primitive.ObjectID) (Accomodation, error)
	UpdatePricing(accomodation Accomodation) error
}
