package auth

import (
	. "auth_service/auth/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type IUserRepository interface {
	Create(user UserModel) (primitive.ObjectID, error)
	FindById(id primitive.ObjectID) (UserModel, error)
	FindByEmail(email string) (UserModel, error)
}
