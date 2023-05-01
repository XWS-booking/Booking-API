package reservation

import (
	. "context"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	. "reservation_service/proto/reservation"
	"reservation_service/reservation/model"
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
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	id := reservationController.ReservationService.Create(model.NewReservation(req))

	return &ReservationId{Id: id.String()}, nil
}

func (reservationController *ReservationController) Delete(ctx Context, req *ReservationId) (*DeleteReservationResponse, error) {
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

func (reservationController *ReservationController) FindAllReservedAccommodations(ctx Context, req *FindAllReservedAccommodationsRequest) (*FindAllReservedAccommodationsResponse, error) {
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
