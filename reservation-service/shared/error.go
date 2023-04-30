package shared

type Error struct {
	Message string
}

func ReservationNotDeleted() *Error {
	return &Error{Message: "Reservation can't be deleted!"}
}

func ReservationNotFound() *Error {
	return &Error{Message: "Reservation not found!"}
}

func ReservationsNotFound() *Error {
	return &Error{Message: "Error when filtering reservations!"}
}

func CheckActiveReservationsError() *Error {
	return &Error{Message: "Error when checking guest's active reservations!"}
}
