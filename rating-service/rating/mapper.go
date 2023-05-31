package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "rating_service/proto/rating"
	. "rating_service/rating/model"
)

func AccommodationRatingFromRateAccommodationRequest(req *RateAccommodationRequest) AccommodationRating {
	accommodationId, _ := primitive.ObjectIDFromHex(req.AccommodationId)
	guestId, _ := primitive.ObjectIDFromHex(req.GuestId)
	return AccommodationRating{
		AccommodationId: accommodationId,
		GuestId:         guestId,
		Rating:          req.Rating,
	}
}
