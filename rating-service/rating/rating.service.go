package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "rating_service/rating/model"
	"rating_service/shared"
	"time"
)

type RatingService struct {
	RatingRepository IRatingRepository
}

func (ratingService *RatingService) CreateAccommdationRating(rating AccommodationRating) primitive.ObjectID {
	rating.Time = time.Now()
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
	rating.Time = time.Now()
	err := ratingService.RatingRepository.UpdateAccommodationRating(rating)
	if err != nil {
		return shared.RatingNotUpdated()
	}
	return nil
}

func (ratingService *RatingService) GetAllAccommodationRatings(id primitive.ObjectID) ([]AccommodationRating, *shared.Error) {
	ratings, err := ratingService.RatingRepository.GetAllByAccommodationId(id)
	if err != nil {
		return nil, shared.ErrorFilteringRatings()
	}
	return ratings, nil
}

func (ratingService *RatingService) GetAverageAccommodationRating(id primitive.ObjectID) (float64, *shared.Error) {
	ratings, err := ratingService.RatingRepository.GetAllByAccommodationId(id)
	if err != nil {
		return -1, shared.ErrorFilteringRatings()
	}
	var avg = 0.0
	var sum = 0.0
	if len(ratings) != 0 {
		for _, r := range ratings {
			sum += float64(r.Rating)
		}
		avg = sum / float64(len(ratings))
	}
	return avg, nil
}

func (ratingService *RatingService) FindAccommodationRatingById(id primitive.ObjectID) (AccommodationRating, *shared.Error) {
	rating, err := ratingService.RatingRepository.FindAccommodationRatingById(id)
	if err != nil {
		return rating, shared.AccommodationRatingNotFound()
	}
	return rating, nil
}
