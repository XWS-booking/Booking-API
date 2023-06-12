package reservation

import (
	. "context"
	"fmt"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "reservation_service/opentelementry"
	. "reservation_service/proto/reservation"
	"reservation_service/reservation/model"
	"reservation_service/shared"
)

func NewReservationController(reservationService *ReservationService) *ReservationController {
	controller := &ReservationController{ReservationService: reservationService}
	return controller
}

type ReservationController struct {
	UnimplementedReservationServiceServer
	ReservationService *ReservationService
}

func (reservationController *ReservationController) Create(ctx Context, req *CreateReservationRequest) (*ReservationId, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "create")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	id := reservationController.ReservationService.Create(model.NewReservation(req))

	return &ReservationId{Id: id.String()}, nil
}

func (reservationController *ReservationController) Delete(ctx Context, req *ReservationId) (*DeleteReservationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "delete")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	reservationId, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	e := reservationController.ReservationService.Delete(reservationId)
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	return &DeleteReservationResponse{Message: "success"}, nil
}

func (reservationController *ReservationController) Confirm(ctx Context, req *ReservationId) (*ReservationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "confirm")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	reservationId, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	reservation, e := reservationController.ReservationService.ConfirmReservation(reservationId)
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	return NewReservationResponse(reservation), nil
}

func (reservationController *ReservationController) Reject(ctx Context, req *ReservationId) (*ReservationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "reject")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	reservationId, err := primitive.ObjectIDFromHex(req.GetId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	reservation, e := reservationController.ReservationService.RejectReservation(reservationId)
	if e != nil {
		return nil, status.Error(codes.Internal, e.Message)
	}
	return NewReservationResponse(reservation), nil
}

func (reservationController *ReservationController) FindAllReservedAccommodations(ctx Context, req *FindAllReservedAccommodationsRequest) (*FindAllReservedAccommodationsResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findAllReservedAccommodations")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	startDate, _ := ptypes.Timestamp(req.StartDate)
	endDate, _ := ptypes.Timestamp(req.EndDate)
	reservedAccommodations, err := reservationController.ReservationService.FindAllReservedAccommodations(startDate, endDate)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Message)
	}
	return &FindAllReservedAccommodationsResponse{Ids: reservedAccommodations}, nil
}

func (reservationController *ReservationController) CheckActiveReservationsForGuest(ctx Context, req *CheckActiveReservationsForGuestRequest) (*CheckActiveReservationsForGuestResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "checkActiveReservationsForGuest")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	id, err := primitive.ObjectIDFromHex(req.GetGuestId())
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	activeReservations, e := reservationController.ReservationService.CheckActiveReservationsForGuest(id)
	if e != nil {
		return &CheckActiveReservationsForGuestResponse{}, status.Error(codes.Internal, e.Message)
	}
	return &CheckActiveReservationsForGuestResponse{
		ActiveReservations: activeReservations,
	}, nil
}

func (reservationController *ReservationController) CheckActiveReservationsForAccommodations(ctx Context, req *CheckActiveReservationsForAccommodationsRequest) (*CheckActiveReservationsForAccommodationsResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "checkActiveReservationsForAccommodations")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	activeReservations, e := reservationController.ReservationService.CheckActiveReservationsForAccommodations(req.GetIds())
	if e != nil {
		return &CheckActiveReservationsForAccommodationsResponse{}, status.Error(codes.Internal, e.Message)
	}
	return &CheckActiveReservationsForAccommodationsResponse{
		ActiveReservations: activeReservations,
	}, nil
}

func (reservationController *ReservationController) CancelReservation(ctx Context, req *CancelReservationRequest) (*ReservationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "cancelReservation")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	reservation, e := reservationController.ReservationService.CancelReservation(shared.StringToObjectId(req.ReservationId))
	if e != nil {
		return &ReservationResponse{}, status.Error(codes.Aborted, e.Message)
	}
	return NewReservationResponse(reservation), nil
}

func (reservationController *ReservationController) IsAccommodationAvailable(ctx Context, req *IsAccommodationAvailableRequest) (*IsAccommodationAvailableResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "isAccommodationAvailable")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	startDate, _ := ptypes.Timestamp(req.StartDate)
	endDate, _ := ptypes.Timestamp(req.EndDate)
	available, e := reservationController.ReservationService.IsAccommodationAvailable(shared.StringToObjectId(req.AccommodationId), startDate, endDate)
	if e != nil {
		return &IsAccommodationAvailableResponse{Available: available}, status.Error(codes.Internal, e.Message)
	}
	return &IsAccommodationAvailableResponse{Available: available}, nil
}

func (reservationController *ReservationController) FindAllByBuyerId(ctx Context, req *FindAllReservationsByBuyerIdRequest) (*FindAllReservationsByBuyerIdResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findAllByBuyerId")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	reservations, e := reservationController.ReservationService.FindAllByBuyerId(shared.StringToObjectId(req.BuyerId))
	if e != nil {
		return &FindAllReservationsByBuyerIdResponse{}, status.Error(codes.Internal, e.Message)
	}
	var reservationResponses []*ReservationResponse
	for _, r := range reservations {
		reservationResponses = append(reservationResponses, NewReservationResponse(r))
	}
	return &FindAllReservationsByBuyerIdResponse{Reservations: reservationResponses}, nil
}

func (reservationController *ReservationController) FindNumberOfBuyersCancellations(ctx Context, req *NumberOfCancellationRequest) (*NumberOfCancellationResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findNumberOfBuyersCancellations")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	numberOfCancellations, e := reservationController.ReservationService.FindNumberOfBuyersCancellations(shared.StringToObjectId(req.BuyerId))
	if e != nil {
		return &NumberOfCancellationResponse{}, status.Error(codes.Internal, e.Message)
	}
	return &NumberOfCancellationResponse{CancellationNumber: int32(numberOfCancellations)}, nil
}

func (reservationController *ReservationController) FindAllByAccommodationId(ctx Context, req *FindAllReservationsByAccommodationIdRequest) (*FindAllReservationsByAccommodationIdResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "findAllByAccommodationId")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	reservations, e := reservationController.ReservationService.FindAllByAccommodationId(shared.StringToObjectId(req.AccommodationId))
	if e != nil {
		return &FindAllReservationsByAccommodationIdResponse{}, status.Error(codes.Internal, e.Message)
	}
	var reservationResponses []*ReservationResponse
	for _, r := range reservations {
		reservationResponses = append(reservationResponses, NewReservationResponse(r))
	}
	return &FindAllReservationsByAccommodationIdResponse{Reservations: reservationResponses}, nil
}

func (reservationController *ReservationController) UpdateReservationRating(ctx Context, req *UpdateReservationRatingRequest) (*UpdateReservationRatingResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "updateReservationRating")
	defer func() { span.End() }()
	fmt.Println(req.Id)
	fmt.Println(req.AccommodationRatingId)
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	e := reservationController.ReservationService.UpdateReservationRating(shared.StringToObjectId(req.Id), shared.StringToObjectId(req.AccommodationRatingId))
	if e != nil {
		return &UpdateReservationRatingResponse{}, status.Error(codes.Internal, e.Message)
	}
	return &UpdateReservationRatingResponse{}, nil
}

func (reservationController *ReservationController) CheckIfGuestHasReservationInAccommodations(ctx Context, req *CheckIfGuestHasReservationInAccommodationsRequest) (*CheckIfGuestHasReservationInAccommodationsResponse, error) {
	_, span := Tp.Tracer(ServiceName).Start(ctx, "checkIfGuestHasReservationInAccommodations")
	defer func() { span.End() }()
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	res, e := reservationController.ReservationService.CheckIfGuestHasReservationInAccommodations(shared.StringToObjectId(req.GuestId), req.AccommodationIds)
	if e != nil {
		return &CheckIfGuestHasReservationInAccommodationsResponse{Res: false}, status.Error(codes.Internal, e.Message)
	}
	return &CheckIfGuestHasReservationInAccommodationsResponse{Res: res}, nil
}
