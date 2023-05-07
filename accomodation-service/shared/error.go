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

func AccomodationNotCreated() *Error {
	return &Error{Message: "Accomodation creation failed!"}
}

func NoMatchingPricingInterval() *Error {
	return &Error{Message: "No matching pricing interval for given time interval!"}
}

func NoPricingIntervalsFound() *Error {
	return &Error{Message: "No pricing intervals found!"}
}

func PricingIntervalsOverlapping() *Error {
	return &Error{Message: "Pricing intervals should not overlap!"}
}
