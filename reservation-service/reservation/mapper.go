package reservation

import (
	"github.com/golang/protobuf/ptypes"
	reservationGrpc "reservation_service/proto/reservation"
	"reservation_service/reservation/model"
)

func NewReservationResponse(reservation model.Reservation) *reservationGrpc.ReservationResponse {
	startDate, _ := ptypes.TimestampProto(reservation.StartDate)
	endDate, _ := ptypes.TimestampProto(reservation.EndDate)
	return &reservationGrpc.ReservationResponse{
		Id:                    reservation.Id.Hex(),
		AccommodationId:       reservation.AccommodationId.Hex(),
		BuyerId:               reservation.BuyerId.Hex(),
		StartDate:             startDate,
		EndDate:               endDate,
		Guests:                reservation.Guests,
		Status:                int32(reservation.Status),
		AccommodationRatingId: reservation.AccommodationRatingId.Hex(),
	}
}
