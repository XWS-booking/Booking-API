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

func (reservationService *ReservationService) Delete(id primitive.ObjectID) (*primitive.ObjectID, *shared.Error) {
	_, e := reservationService.ReservationRepository.FindById(id)
	if e != nil {
		return nil, shared.ReservationNotFound()
	}
	accomId, error := reservationService.ReservationRepository.Delete(id)
	if error != nil {
		return nil, shared.ReservationNotDeleted()
	}
	return accomId, nil
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

func (reservationService *ReservationService) CancelReservation(reservationId primitive.ObjectID) (Reservation, *shared.Error) {
	reservation, err := reservationService.ReservationRepository.FindById(reservationId)
	if err != nil {
		return reservation, shared.ReservationCancelationFailed()
	}
	e := reservation.Cancel()
	if e != nil {
		return reservation, e
	}
	reservationService.ReservationRepository.UpdateReservation(reservation)
	return reservation, nil
}

func (reservationService *ReservationService) ConfirmReservation(id primitive.ObjectID) (Reservation, *shared.Error) {
	reservation, err := reservationService.ReservationRepository.FindById(id)
	if err != nil {
		return reservation, shared.ReservationConfirmationFailed()
	}
	e := reservation.Confirm()
	if e != nil {
		return reservation, e
	}
	err = reservationService.ReservationRepository.UpdateReservation(reservation)
	if err != nil {
		return reservation, shared.ReservationUpdateFailed()
	}
	e = reservationService.CancelOverlappingReservations(reservation)
	if e != nil {
		return reservation, e
	}
	return reservation, nil
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

func (reservationService *ReservationService) RejectReservation(id primitive.ObjectID) (Reservation, *shared.Error) {
	reservation, err := reservationService.ReservationRepository.FindById(id)
	if err != nil {
		return reservation, shared.ReservationRejectionFailed()
	}
	e := reservation.Reject()
	if e != nil {
		return reservation, e
	}
	reservationService.ReservationRepository.UpdateReservation(reservation)
	return reservation, nil
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

func (reservationService *ReservationService) FindNumberOfBuyersCancellations(id primitive.ObjectID) (int, *shared.Error) {
	numberOfCancellations, err := reservationService.ReservationRepository.FindNumberOfBuyersCancellations(id)
	if err != nil {
		return numberOfCancellations, shared.ReservationsNotFound()
	}
	return numberOfCancellations, nil
}

func (reservationService *ReservationService) FindAllByAccommodationId(id primitive.ObjectID) ([]Reservation, *shared.Error) {
	reservations, err := reservationService.ReservationRepository.FindAllByAccommodationId(id)
	if err != nil {
		return reservations, shared.ReservationsNotFound()
	}
	return reservations, nil
}

func (reservationService *ReservationService) UpdateReservationRating(id primitive.ObjectID, accommodationRatingId primitive.ObjectID) *shared.Error {
	reservation, err := reservationService.ReservationRepository.FindById(id)
	if err != nil {
		return shared.ReservationNotFound()
	}
	reservation.AccommodationRatingId = accommodationRatingId
	err = reservationService.ReservationRepository.UpdateReservation(reservation)
	if err != nil {
		return shared.ReservationUpdateFailed()
	}
	return nil
}

func (reservationService *ReservationService) CheckIfGuestHasReservationInAccommodations(guestId primitive.ObjectID, accommodationIds []string) (bool, *shared.Error) {
	for _, id := range accommodationIds {
		resId, err := reservationService.ReservationRepository.CheckIfGuestHasReservationInAccommodation(guestId, shared.StringToObjectId(id))
		if err != nil {
			return false, shared.SomethingWentWrongWhenFindingReservation()
		}
		if resId {
			return resId, nil
		}
	}
	return false, nil
}
