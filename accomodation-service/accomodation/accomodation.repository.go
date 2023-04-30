package accomodation

import (
	. "accomodation_service/accomodation/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type AccomodationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (accomodationRepository *AccomodationRepository) FindAll(city string, guests int32) ([]Accomodation, error) {
	collection := accomodationRepository.getCollection("accomodations")
	var accommodations []Accomodation

	filter := bson.D{{Key: "city", Value: bson.D{{Key: "$regex", Value: "(?i).*" + city + ".*"}}}}
	if guests != -1 {
		filter = bson.D{{Key: "city", Value: bson.D{{Key: "$regex", Value: "(?i).*" + city + ".*"}}},
			{Key: "min_guests", Value: bson.D{{Key: "$lte", Value: guests}}},
			{Key: "max_guests", Value: bson.D{{Key: "$gte", Value: guests}}}}
	}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return accommodations, err
	}

	for cur.Next(context.TODO()) {
		var elem Accomodation
		err := cur.Decode(&elem)
		if err != nil {
			return accommodations, err
		}
		accommodations = append(accommodations, elem)
	}
	return accommodations, nil
}

func (accomodationRepository *AccomodationRepository) FindAllByOwnerId(id primitive.ObjectID) ([]Accomodation, error) {
	collection := accomodationRepository.getCollection("accomodations")
	var accommodations []Accomodation

	filter := bson.M{"owner_id": id}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return accommodations, err
	}

	for cur.Next(context.TODO()) {
		var elem Accomodation
		err := cur.Decode(&elem)
		if err != nil {
			return accommodations, err
		}
		accommodations = append(accommodations, elem)
	}
	return accommodations, nil
}

func (accommodationRepository *AccomodationRepository) DeleteByOwnerId(id primitive.ObjectID) error {
	collection := accommodationRepository.getCollection("accomodations")
	filter := bson.M{"owner_id": id}
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (accomodationRepository *AccomodationRepository) getCollection(key string) *mongo.Collection {
	return accomodationRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
