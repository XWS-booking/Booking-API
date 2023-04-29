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
