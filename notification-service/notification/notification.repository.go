package notification

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type NotificationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (ratingRepository *NotificationRepository) getCollection(key string) *mongo.Collection {
	return ratingRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
