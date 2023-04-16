package model

import (
	authGrpc "auth_service/proto/auth"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole int64

const (
	GUEST             UserRole = 0
	HOST              UserRole = 1
	NOT_AUTHENTICATED UserRole = 2
)

type User struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name     string             `bson:"name" json:"name"`
	Surname  string             `bson:"surname" json:"surname"`
	Email    string             `bson:"email" json:"email"`
	Password string             `bson:"password" json:"password"`
	Role     UserRole           `bson:"role" json:"role"`
}

func (user *User) MapFromProto(protoUser *authGrpc.User) {
	user.Name = protoUser.Name
	user.Surname = protoUser.Surname
	user.Email = protoUser.Email
	user.Password = protoUser.Password
	user.Role = UserRole(protoUser.Role)
}
