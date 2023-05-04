package reservation

import (
	. "context"
	"github.com/golang/protobuf/ptypes"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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

func (reservationController *ReservationController) CancelReservation(ctx Context, req *CancelReservationRequest) (*CancelReservationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}
	e := reservationController.ReservationService.CancelReservation(shared.StringToObjectId(req.ReservationId))
	if e != nil {
		return &CancelReservationResponse{}, status.Error(codes.Aborted, e.Message)
	}
	return &CancelReservationResponse{}, nil
}

func (reservationController *ReservationController) IsAccommodationAvailable(ctx Context, req *IsAccommodationAvailableRequest) (*IsAccommodationAvailableResponse, error) {
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
