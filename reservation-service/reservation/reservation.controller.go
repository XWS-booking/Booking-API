package reservation

import (
	. "context"
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
	reservationId, _ := primitive.ObjectIDFromHex(req.GetId())
	err := reservationController.ReservationService.Delete(reservationId)
	if err != nil {
		return nil, status.Error(codes.Unknown, err.Message)
	}
	return &DeleteReservationResponse{Message: "success"}, nil
}
