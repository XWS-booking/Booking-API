package shared

type Error struct {
	Message string
}

func RatingNotDeleted() *Error {
	return &Error{Message: "Rating can't be deleted!"}
}

func RatingNotUpdated() *Error {
	return &Error{Message: "Rating can't be updated!"}
}
