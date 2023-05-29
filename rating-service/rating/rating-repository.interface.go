package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "rating_service/rating/model"
)

type IRatingRepository interface {
	CreateAccommodationRating(rating AccommodationRating) (primitive.ObjectID, error)
	DeleteAccommodationRating(id primitive.ObjectID) error
	UpdateAccommodationRating(rating AccommodationRating) error
}
