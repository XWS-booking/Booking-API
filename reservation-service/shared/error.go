package shared

type Error struct {
	Message string
}

func ReservationNotDeleted() *Error {
	return &Error{Message: "Reservation can't be deleted!"}
}

func ReservationNotConfirmed() *Error {
	return &Error{Message: "Reservation can't be confirmed!"}
}

func ReservationNotFound() *Error {
	return &Error{Message: "Reservation not found!"}
}

func ReservationsNotFound() *Error {
	return &Error{Message: "Error when filtering reservations!"}
}

func CheckActiveReservationsError() *Error {
	return &Error{Message: "Error when checking active reservations!"}
}

func ReservationCancelationFailed() *Error {
	return &Error{Message: "Error when canceling reservation!"}
}

func ReservationConfirmationFailed() *Error {
	return &Error{Message: "Error when confirming reservation!"}
}

func ReservationRejectionFailed() *Error {
	return &Error{Message: "Error when rejecting reservation!"}
}

func ReservationCancelationTooLate() *Error {
	return &Error{Message: "Can't cancel later than day before reservation start!"}
}
