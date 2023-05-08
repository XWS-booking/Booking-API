package reservation

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	. "reservation_service/reservation/model"
	"reservation_service/shared"
	"time"
)

type ReservationService struct {
	ReservationRepository IReservationRepository
}

func (reservationService *ReservationService) Create(reservation Reservation) primitive.ObjectID {
	reservation.Status = 0
	created, error := reservationService.ReservationRepository.Create(reservation)
	if error != nil {
		return created
	}
	return created
}

func (reservationService *ReservationService) Delete(id primitive.ObjectID) *shared.Error {
	_, e := reservationService.ReservationRepository.FindById(id)
	if e != nil {
		return shared.ReservationNotFound()
	}
	error := reservationService.ReservationRepository.Delete(id)
	if error != nil {
		return shared.ReservationNotDeleted()
	}
	return nil
}

func (reservationService ReservationService) FindAllReservedAccommodations(startDate time.Time, endDate time.Time) ([]string, *shared.Error) {
	accommodationIds, err := reservationService.ReservationRepository.FindAllReservedAccommodations(startDate, endDate)
	if err != nil {
		return accommodationIds, shared.ReservationsNotFound()
	}
	return accommodationIds, nil
}

func (reservationService *ReservationService) CheckActiveReservationsForGuest(id primitive.ObjectID) (bool, *shared.Error) {
	activeReservations, err := reservationService.ReservationRepository.CheckActiveReservationsForGuest(id)
	if err != nil {
		return activeReservations, shared.CheckActiveReservationsError()
	}
	if !activeReservations {
		err = reservationService.ReservationRepository.DeleteByBuyerId(id)
		if err != nil {
			return activeReservations, shared.ReservationNotDeleted()
		}
	}
	return activeReservations, nil
}

func (reservationService *ReservationService) CheckActiveReservationsForAccommodations(accommodationIds []string) (bool, *shared.Error) {
	for _, idStr := range accommodationIds {
		id, _ := primitive.ObjectIDFromHex(idStr)
		activeReservations, err := reservationService.ReservationRepository.CheckActiveReservationsForAccommodation(id)
		if err != nil {
			return false, shared.CheckActiveReservationsError()
		}
		if activeReservations {
			return true, nil
		}
	}
	for _, idStr := range accommodationIds {
		id, _ := primitive.ObjectIDFromHex(idStr)
		err := reservationService.ReservationRepository.DeleteByAccommodationId(id)
		if err != nil {
			return false, shared.ReservationNotDeleted()
		}
	}
	return false, nil
}

func (reservationService *ReservationService) CancelReservation(reservationId primitive.ObjectID) *shared.Error {
	reservation, err := reservationService.ReservationRepository.FindById(reservationId)
	if err != nil {
		return shared.ReservationCancelationFailed()
	}
	e := reservation.Cancel()
	if e != nil {
		return e
	}
	reservationService.ReservationRepository.UpdateReservation(reservation)
	return nil
}

func (reservationService *ReservationService) ConfirmReservation(id primitive.ObjectID) *shared.Error {
	reservation, err := reservationService.ReservationRepository.FindById(id)
	if err != nil {
		return shared.ReservationConfirmationFailed()
	}
	e := reservation.Confirm()
	if e != nil {
		return e
	}
	err = reservationService.ReservationRepository.UpdateReservation(reservation)
	if err != nil {
		return shared.ReservationUpdateFailed()
	}
	e = reservationService.CancelOverlappingReservations(reservation)
	if e != nil {
		return e
	}
	return nil
}

func (reservationService *ReservationService) CancelOverlappingReservations(reservation Reservation) *shared.Error {
	id2, _ := primitive.ObjectIDFromHex(reservation.AccommodationId.Hex())
	reservations, err := reservationService.ReservationRepository.FindAllPendingByAccommodationId(id2)
	if err != nil {
		return shared.ReservationNotFound()
	}
	for _, reservationToCheck := range reservations {
		if reservationToCheck.IsOverlapping(reservation) {
			reservationService.RejectReservation(reservationToCheck.Id)
		}
	}
	return nil
}

func (reservationService *ReservationService) RejectReservation(id primitive.ObjectID) *shared.Error {
	reservation, err := reservationService.ReservationRepository.FindById(id)
	if err != nil {
		return shared.ReservationRejectionFailed()
	}
	e := reservation.Reject()
	if e != nil {
		return e
	}
	reservationService.ReservationRepository.UpdateReservation(reservation)
	return nil
}

func (reservationService *ReservationService) IsAccommodationAvailable(id primitive.ObjectID, startDate time.Time, endDate time.Time) (bool, *shared.Error) {
	available, err := reservationService.ReservationRepository.IsAccommodationAvailable(id, startDate, endDate)
	if err != nil {
		return available, shared.ReservationsNotFound()
	}
	return available, nil
}

func (reservationService *ReservationService) FindAllByBuyerId(id primitive.ObjectID) ([]Reservation, *shared.Error) {
	reservations, err := reservationService.ReservationRepository.FindAllByBuyerId(id)
	if err != nil {
		return reservations, shared.ReservationsNotFound()
	}
	return reservations, nil
}

func (reservationService *ReservationService) FindAllByAccommodationId(id primitive.ObjectID) ([]Reservation, *shared.Error) {
	reservations, err := reservationService.ReservationRepository.FindAllByAccommodationId(id)
	if err != nil {
		return reservations, shared.ReservationsNotFound()
	}
	return reservations, nil
}
