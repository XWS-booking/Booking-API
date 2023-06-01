package rating

import (
	. "context"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "rating_service/proto/rating"
	"rating_service/shared"
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
	return &RateAccommodationResponse{Id: id.Hex()}, nil
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
	e := ratingController.RatingService.UpdateAccommodationRating(shared.StringToObjectId(req.Id), req.Rating)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}
	return &UpdateAccommodationRatingResponse{}, nil
}

func (ratingController *RatingController) GetAllAccommodationRatings(ctx Context, req *GetAllAccommodationRatingsRequest) (*GetAllAccommodationRatingsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetAccommodationId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	ratings, e := ratingController.RatingService.GetAllAccommodationRatings(id)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}
	var ratingResponses []*AccommodationRatingItem
	for _, r := range ratings {
		time, _ := ptypes.TimestampProto(r.Time)
		ratingResponses = append(ratingResponses, &AccommodationRatingItem{
			Id:              r.Id.Hex(),
			AccommodationId: r.AccommodationId.Hex(),
			GuestId:         r.GuestId.Hex(),
			Rating:          r.Rating,
			Time:            time,
		})
	}
	return &GetAllAccommodationRatingsResponse{Ratings: ratingResponses}, nil
}

func (ratingController *RatingController) GetAverageAccommodationRating(ctx Context, req *GetAverageAccommodationRatingRequest) (*GetAverageAccommodationRatingResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetAccommodationId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	rating, e := ratingController.RatingService.GetAverageAccommodationRating(id)
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}

	return &GetAverageAccommodationRatingResponse{Rating: rating}, nil
}

func (ratingController *RatingController) FindAccommodationRatingById(ctx Context, req *FindAccommodationRatingByIdRequest) (*FindAccommodationRatingByIdResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	rating, e := ratingController.RatingService.FindAccommodationRatingById(shared.StringToObjectId(req.Id))
	if e != nil {
		return nil, status.Error(codes.Aborted, e.Message)
	}

	return &FindAccommodationRatingByIdResponse{
		Id:              rating.Id.Hex(),
		AccommodationId: rating.AccommodationId.Hex(),
		GuestId:         rating.GuestId.Hex(),
		Rating:          rating.Rating,
	}, nil
}

func (ratingController *RatingController) RateHost(ctx Context, req *RateHostRequest) (*RateHostResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	res, err := ratingController.RatingService.RateHost(HostRatingFromRateHostRequest(req))
	if err != nil {
		return nil, status.Error(codes.Aborted, err.Message)
	}

	return &RateHostResponse{Id: res.Hex()}, nil
}
