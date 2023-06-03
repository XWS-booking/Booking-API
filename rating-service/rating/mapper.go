package rating

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "rating_service/proto/rating"
	. "rating_service/rating/model"
	"rating_service/shared"
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

func HostRatingFromRateHostRequest(req *RateHostRequest) HostRating {
	hostId, _ := primitive.ObjectIDFromHex(req.HostId)
	guestId, _ := primitive.ObjectIDFromHex(req.GuestId)
	return HostRating{
		HostId:  hostId,
		GuestId: guestId,
		Rating:  req.Rating,
	}
}

func UpdateHostRatingFromUpdateRateHostRequest(req *UpdateHostRatingRequest) HostRating {

	return HostRating{
		Id:     shared.StringToObjectId(req.Id),
		Rating: req.Rating,
	}
}
