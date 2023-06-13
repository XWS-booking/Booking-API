package rating

import (
	"fmt"
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

func (ratingService *RatingService) DeleteAccommodationRating(id primitive.ObjectID) error {
	e := ratingService.RatingRepository.DeleteAccommodationRating(id)
	if e != nil {
		return e
	}
	return nil
}

func (ratingService *RatingService) UpdateAccommodationRating(id primitive.ObjectID, rating int32) error {
	res, err := ratingService.RatingRepository.FindAccommodationRatingById(id)
	if err != nil {
		return err
	}
	res.Time = time.Now()
	res.Rating = rating
	err = ratingService.RatingRepository.UpdateAccommodationRating(res)
	if err != nil {
		return err
	}
	return nil
}

func (ratingService *RatingService) GetAllAccommodationRatings(id primitive.ObjectID) ([]AccommodationRating, error) {
	ratings, err := ratingService.RatingRepository.GetAllByAccommodationId(id)
	if err != nil {
		return nil, err
	}
	return ratings, nil
}

func (ratingService *RatingService) GetAverageAccommodationRating(id primitive.ObjectID) (float64, error) {
	ratings, err := ratingService.RatingRepository.GetAllByAccommodationId(id)
	if err != nil {
		return -1, err
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

func (ratingService *RatingService) FindAccommodationRatingById(id primitive.ObjectID) (AccommodationRating, error) {
	rating, err := ratingService.RatingRepository.FindAccommodationRatingById(id)
	if err != nil {
		return rating, err
	}
	return rating, nil
}

func (ratingService *RatingService) RateHost(hostRating HostRating) (primitive.ObjectID, error) {
	hostRating.Time = time.Now()
	res, err := ratingService.RatingRepository.CreateHostRating(hostRating)
	if err != nil {
		return res, err
	}
	return res, nil
}

func (ratingService *RatingService) UpdateHostRating(rating HostRating) (HostRating, error) {
	currentRating, err := ratingService.RatingRepository.FindHostRatingById(rating.Id)
	if err != nil {
		return HostRating{}, err
	}
	currentRating.Rating = rating.Rating
	currentRating.Time = time.Now()
	res, err2 := ratingService.RatingRepository.UpdateHostRating(currentRating)
	if err2 != nil {
		return HostRating{}, err
	}
	return res, nil
}

func (ratingService *RatingService) DeleteHostRating(id string) (*primitive.ObjectID, error) {
	hostId, err := ratingService.RatingRepository.DeleteHostRating(shared.StringToObjectId(id))
	if err != nil {
		return nil, err
	}
	return hostId, nil
}

func (ratingService *RatingService) GetHostRatings(id string) ([]HostRating, error) {
	ratings, err := ratingService.RatingRepository.GetHostRatings(shared.StringToObjectId(id))
	if err != nil {
		return []HostRating{}, err
	}
	return ratings, nil
}

func (ratingService *RatingService) CalculateHostAverageRate(ratings []HostRating) float32 {
	var averageRate float32 = 0
	for _, rating := range ratings {
		averageRate += float32(rating.Rating)
	}
	fmt.Println(averageRate / float32(len(ratings)))
	return averageRate / float32(len(ratings))
}
