package reservation

import (
	. "context"
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

func (reservationController *ReservationController) Create(ctx Context, req *CreateReservationRequest) (*CreateReservationResponse, error) {
	if req == nil {
		return nil, status.Error(codes.Aborted, "Something wrong with data")
	}

	id := reservationController.ReservationService.Create(model.NewReservation(req))

	return &CreateReservationResponse{Id: id.String()}, nil
}
