package shared

type Error struct {
	Message string
}

func AccommodationNotDeleted() *Error {
	return &Error{Message: "Accommodation can't be deleted!"}
}

func AccommodationsNotFound() *Error {
	return &Error{Message: "Error when filtering accommodations!"}
}
