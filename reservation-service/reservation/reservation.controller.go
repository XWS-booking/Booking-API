package reservation

import (
	. "reservation_service/proto/reservation"
)

func NewReservationController(reservationService *ReservationService) *ReservationController {
	controller := &ReservationController{ReservationService: reservationService}
	return controller
}

type ReservationController struct {
	UnimplementedReservationServiceServer
	ReservationService *ReservationService
}
