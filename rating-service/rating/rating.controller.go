package rating

import (
	. "context"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"net/http"
	. "rating_service/opentelementry"
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
	_, span := Tp.Tracer(ServiceName).Start(ctx, "rateAccommodation")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id := ratingController.RatingService.CreateAccommdationRating(AccommodationRatingFromRateAccommodationRequest(req))
	return &RateAccommodationResponse{Id: id.Hex()}, nil
}

func (ratingController *RatingController) DeleteAccommodationRating(ctx Context, req *DeleteAccommodationRatingRequest) (*DeleteAccommodationRatingResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "deleteAccommodationRating")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		HttpError(err, span, http.StatusBadRequest)
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	e := ratingController.RatingService.DeleteAccommodationRating(id)
	if e != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, e.Error())
	}
	return &DeleteAccommodationRatingResponse{}, nil
}

func (ratingController *RatingController) UpdateAccommodationRating(ctx Context, req *UpdateAccommodationRatingRequest) (*UpdateAccommodationRatingResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "updateAccommodationRating")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	e := ratingController.RatingService.UpdateAccommodationRating(shared.StringToObjectId(req.Id), req.Rating)
	if e != nil {
		HttpError(e, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, e.Error())
	}
	return &UpdateAccommodationRatingResponse{}, nil
}

func (ratingController *RatingController) GetAllAccommodationRatings(ctx Context, req *GetAllAccommodationRatingsRequest) (*GetAllAccommodationRatingsResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "getAllAccommodationRatings")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetAccommodationId())
	if err != nil {
		HttpError(err, span, http.StatusBadRequest)
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	ratings, e := ratingController.RatingService.GetAllAccommodationRatings(id)
	if e != nil {
		HttpError(e, span, http.StatusInternalServerError)
		return nil, status.Error(codes.Aborted, e.Error())
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
	_, span := Tp.Tracer(ServiceName).Start(ctx, "getAverageAccommodationRating")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetAccommodationId())
	if err != nil {
		HttpError(err, span, http.StatusBadRequest)
		return nil, status.Error(http.StatusBadRequest, err.Error())
	}
	rating, e := ratingController.RatingService.GetAverageAccommodationRating(id)
	if e != nil {
		HttpError(e, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, e.Error())
	}

	return &GetAverageAccommodationRatingResponse{Rating: rating}, nil
}

func (ratingController *RatingController) FindAccommodationRatingById(ctx Context, req *FindAccommodationRatingByIdRequest) (*FindAccommodationRatingByIdResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findAccommodationRatingById")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	rating, e := ratingController.RatingService.FindAccommodationRatingById(shared.StringToObjectId(req.Id))
	if e != nil {
		HttpError(e, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, e.Error())
	}

	return &FindAccommodationRatingByIdResponse{
		Id:              rating.Id.Hex(),
		AccommodationId: rating.AccommodationId.Hex(),
		GuestId:         rating.GuestId.Hex(),
		Rating:          rating.Rating,
	}, nil
}

func (ratingController *RatingController) RateHost(ctx Context, req *RateHostRequest) (*RateHostResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "rateHost")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	res, err := ratingController.RatingService.RateHost(HostRatingFromRateHostRequest(req))
	if err != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &RateHostResponse{Id: res.Hex()}, nil
}

func (ratingController *RatingController) UpdateHostRating(ctx Context, req *UpdateHostRatingRequest) (*UpdateHostRatingResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "updateHostRating")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	res, err := ratingController.RatingService.UpdateHostRating(UpdateHostRatingFromUpdateRateHostRequest(req))
	if err != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &UpdateHostRatingResponse{Id: res.Id.Hex(), HostId: res.HostId.Hex(), GuestId: res.GuestId.Hex(), Rating: res.Rating}, nil
}

func (ratingController *RatingController) DeleteHostRating(ctx Context, req *DeleteHostRatingRequest) (*DeleteHostRatingResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "deleteHostRating")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	err := ratingController.RatingService.DeleteHostRating(req.Id)
	if err != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return nil, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &DeleteHostRatingResponse{}, nil
}

func (ratingController *RatingController) GetHostRatings(ctx Context, req *GetHostRatingsRequest) (*GetHostRatingsResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "getHostRatings")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	ratings, err := ratingController.RatingService.GetHostRatings(req.HostId)
	averageRate := ratingController.RatingService.CalculateHostAverageRate(ratings)

	var ratingResponses []*HostRatingItem
	for _, r := range ratings {
		time, _ := ptypes.TimestampProto(r.Time)
		ratingResponses = append(ratingResponses, &HostRatingItem{
			Id:      r.Id.Hex(),
			HostId:  r.HostId.Hex(),
			GuestId: r.GuestId.Hex(),
			Rating:  r.Rating,
			Time:    time,
		})
	}
	if err != nil {
		HttpError(err, span, http.StatusInternalServerError)
		return &GetHostRatingsResponse{Ratings: ratingResponses}, status.Error(http.StatusInternalServerError, err.Error())
	}

	return &GetHostRatingsResponse{Ratings: ratingResponses, AverageRate: float64(averageRate)}, nil

}
