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

func ErrorFilteringRatings() *Error {
	return &Error{Message: "Error when filtering ratings!"}
}

func AccommodationRatingNotFound() *Error {
	return &Error{Message: "Accommodation rating not found!"}
}

func UnsuccessfulHostRating() *Error {
	return &Error{Message: "Unsuccessful host rating!!"}
}

func HostRatingNotFound() *Error {
	return &Error{Message: "Host rating not found!"}
}

func ErrorWhenGettingRatings() *Error {
	return &Error{Message: "Error while getting host ratings!"}

}
