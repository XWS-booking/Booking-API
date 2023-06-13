package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRole int64

const (
	GUEST             UserRole = 0
	HOST              UserRole = 1
	NOT_AUTHENTICATED UserRole = 2
)

type DeleteStatus int64

const (
	PENDING DeleteStatus = 0
	DELETED DeleteStatus = 1
	ACTIVE  DeleteStatus = 2
)

type User struct {
	Id            primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name          string             `bson:"name" json:"name"`
	Surname       string             `bson:"surname" json:"surname"`
	Email         string             `bson:"email" json:"email"`
	Password      string             `bson:"password" json:"password"`
	Role          UserRole           `bson:"role" json:"role"`
	Street        string             `bson:"street" json:"street"`
	StreetNumber  string             `bson:"street_number" json:"streetNumber"`
	City          string             `bson:"city" json:"city"`
	ZipCode       string             `bson:"zip_code" json:"zipCode"`
	Country       string             `bson:"country" json:"country"`
	Username      string             `bson:"username" json:"username"`
	DeleteStatus  DeleteStatus       `bson:"delete_status" json:"deleteStatus"`
	Distinguished bool               `bson:"distinguished" json:"distinguished"`
}
