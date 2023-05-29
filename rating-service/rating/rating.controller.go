package rating

import (
	. "context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "rating_service/proto/rating"
)

func NewRatingController(ratingService *RatingService) *RatingController {
	controller := &RatingController{RatingService: ratingService}
	return controller
}

type RatingController struct {
	UnimplementedRatingServiceServer
	RatingService *RatingService
}

func (ratingController *RatingController) RateAccommodation(ctx Context, req *RateAccommodationRequest) (*RateAccommodationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id := ratingController.RatingService.CreateAccommdationRating(AccommodationRatingFromRateAccommodationRequest(req))
	return &RateAccommodationResponse{Id: id.String()}, nil
}

func (ratingController *RatingController) DeleteAccommodationRating(ctx Context, req *DeleteAccommodationRatingRequest) (*DeleteAccommodationRatingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	e := ratingController.RatingService.DeleteAccommodationRating(id)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}
	return &DeleteAccommodationRatingResponse{}, nil
}

func (ratingController *RatingController) UpdateAccommodationRating(ctx Context, req *UpdateAccommodationRatingRequest) (*UpdateAccommodationRatingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	e := ratingController.RatingService.UpdateAccommodationRating(AccommodationRatingFromUpdateAccommodationRatingRequest(req))
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}
	return &UpdateAccommodationRatingResponse{}, nil
}
