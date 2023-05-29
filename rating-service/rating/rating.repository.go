package rating

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	. "rating_service/rating/model"
)

type RatingRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (ratingRepository *RatingRepository) getCollection(key string) *mongo.Collection {
	return ratingRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}

func (ratingRepository *RatingRepository) CreateAccommodationRating(rating AccommodationRating) (primitive.ObjectID, error) {
	collection := ratingRepository.getCollection("accommodation_ratings")
	res, err := collection.InsertOne(context.TODO(), rating)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (reservationRepository *RatingRepository) DeleteAccommodationRating(id primitive.ObjectID) error {
	collection := reservationRepository.getCollection("accommodation_ratings")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (reservationRepository *RatingRepository) UpdateAccommodationRating(rating AccommodationRating) error {
	collection := reservationRepository.getCollection("accommodation_ratings")
	filter := bson.M{"_id": rating.Id}
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": rating})
	if err != nil {
		return err
	}
	return nil
}

func (ratingRepository *RatingRepository) GetAllByAccommodationId(id primitive.ObjectID) ([]AccommodationRating, error) {
	collection := ratingRepository.getCollection("accommodation_ratings")
	var ratings []AccommodationRating
	filter := bson.M{"accommodation_id": id}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return ratings, err
	}

	for cur.Next(context.TODO()) {
		var elem AccommodationRating
		err := cur.Decode(&elem)
		if err != nil {
			return ratings, err
		}
		ratings = append(ratings, elem)
	}
	return ratings, nil
}