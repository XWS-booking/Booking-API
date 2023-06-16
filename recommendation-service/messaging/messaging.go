package messaging

import "recommendation_service/recommendation/model"

type Publisher interface {
	Publish(message interface{}) error
}

type Subscriber interface {
	Subscribe(function interface{}) error
}

type OperationType int32

const (
	CREATE OperationType = 0
	UPDATE OperationType = 1
	DELETE OperationType = 2
)

type AccommodationMessage struct {
	Type          OperationType
	Accommodation model.Accommodation
}

type RatingMessage struct {
	Type   OperationType
	Rating model.Rating
}

type UserMessage struct {
	Type OperationType
	User model.User
}
