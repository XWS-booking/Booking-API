package accomodation

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type AccomodationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}
