package reservation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	. "reservation_service/reservation/model"
)

type ReservationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (reservationRepository *ReservationRepository) Create(reservation Reservation) (primitive.ObjectID, error) {
	collection := reservationRepository.getCollection("reservations")
	res, err := collection.InsertOne(context.TODO(), reservation)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (reservationRepository *ReservationRepository) getCollection(key string) *mongo.Collection {
	return reservationRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
