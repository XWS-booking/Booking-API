package reservation

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type ReservationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}
