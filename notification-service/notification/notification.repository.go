package notification

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	. "notification_service/notification/model"
	"os"
)

type NotificationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (ratingRepository *NotificationRepository) getCollection(key string) *mongo.Collection {
	return ratingRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}

func (reservationRepository *NotificationRepository) Create(notification Notification) (primitive.ObjectID, error) {
	collection := reservationRepository.getCollection("notifications")
	res, err := collection.InsertOne(context.TODO(), notification)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (notificationRepository *NotificationRepository) FindById(id primitive.ObjectID) (Notification, error) {
	collection := notificationRepository.getCollection("notifications")
	var notification Notification
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&notification)
	if err != nil {
		return Notification{}, err
	}
	return notification, nil
}

func (notificationRepository *NotificationRepository) Update(notification Notification) error {
	collection := notificationRepository.getCollection("notifications")
	filter := bson.M{"_id": notification.UserId}
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": notification})
	if err != nil {
		return err
	}
	return nil
}
