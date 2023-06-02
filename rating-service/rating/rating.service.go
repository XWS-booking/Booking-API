package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"math"
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

func (ratingService *RatingService) UpdateAccommodationRating(id primitive.ObjectID, rating int32) *shared.Error {
	res, err := ratingService.RatingRepository.FindAccommodationRatingById(id)
	if err != nil {
		return shared.AccommodationRatingNotFound()
	}
	res.Time = time.Now()
	res.Rating = rating
	err = ratingService.RatingRepository.UpdateAccommodationRating(res)
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
		avg = math.Round(avg*100) / 100
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

func (ratingService *RatingService) RateHost(hostRating HostRating) (primitive.ObjectID, *shared.Error) {
	hostRating.Time = time.Now()
	res, err := ratingService.RatingRepository.CreateHostRating(hostRating)
	if err != nil {
		return res, shared.UnsuccessfulHostRating()
	}
	return res, nil
}

func (ratingService *RatingService) UpdateHostRating(rating HostRating) (HostRating, *shared.Error) {
	currentRating, err := ratingService.RatingRepository.FindHostRatingById(rating.Id)
	if err != nil {
		return HostRating{}, shared.HostRatingNotFound()
	}
	currentRating.Rating = rating.Rating
	currentRating.Time = time.Now()
	res, err2 := ratingService.RatingRepository.UpdateHostRating(currentRating)
	if err2 != nil {
		return HostRating{}, shared.RatingNotUpdated()
	}
	return res, nil
}

func (ratingService *RatingService) DeleteHostRating(id string) *shared.Error {
	err := ratingService.RatingRepository.DeleteHostRating(shared.StringToObjectId(id))
	if err != nil {
		return shared.RatingNotDeleted()
	}
	return nil
}
