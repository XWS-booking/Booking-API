package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "rating_service/rating/model"
	"rating_service/shared"
)

type RatingService struct {
	RatingRepository IRatingRepository
}

func (ratingService *RatingService) CreateAccommdationRating(rating AccommodationRating) primitive.ObjectID {
	created, error := ratingService.RatingRepository.CreateAccommodationRating(rating)
	if error != nil {
		return created
	}
	return created
}

func (ratingService *RatingService) DeleteAccommodationRating(id primitive.ObjectID) *shared.Error {
	e := ratingService.RatingRepository.DeleteAccommodationRating(id)
	if e != nil {
		return shared.RatingNotDeleted()
	}
	return nil
}

func (ratingService *RatingService) UpdateAccommodationRating(rating AccommodationRating) *shared.Error {
	err := ratingService.RatingRepository.UpdateAccommodationRating(rating)
	if err != nil {
		return shared.RatingNotUpdated()
	}
	return nil
}
