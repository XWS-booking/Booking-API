package user

import "recommendation_service/recommendation/model"

type IUserRepository interface {
	Create(user model.User) error
	Delete(user model.User) error
	GetRecommended(user model.User) ([]model.Accommodation, error)
}
