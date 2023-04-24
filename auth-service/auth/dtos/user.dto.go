package dtos

import (
	. "auth_service/auth/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserDto struct {
	Id      primitive.ObjectID `json:"id"`
	Name    string             `json:"name"`
	Surname string             `json:"surname"`
	Email   string             `json:"email"`
	Role    UserRole           `json:"role"`
}

func NewUserDto(user User) *UserDto {
	return &UserDto{
		Id:      user.Id,
		Name:    user.Name,
		Surname: user.Surname,
		Email:   user.Email,
		Role:    user.Role,
	}
}
