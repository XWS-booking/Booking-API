package accomodation

import (
	"accomodation_service/accomodation/model"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type AccomodationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (accomodationRepository *AccomodationRepository) FindAll(city string, guests int32) ([]model.Accomodation, error) {
	collection := accomodationRepository.getCollection("accomodations")
	var accomodations []model.Accomodation
	filter := bson.D{}
	filter = bson.D{{Key: "city", Value: bson.D{{Key: "$regex", Value: "(?i).*" + city + ".*"}}},
		{Key: "min_guests", Value: bson.D{{Key: "$lte", Value: guests}}},
		{Key: "max_guests", Value: bson.D{{Key: "$gte", Value: guests}}}}

	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return accomodations, err
	}

	for cur.Next(context.TODO()) {
		var elem model.Accomodation
		err := cur.Decode(&elem)
		if err != nil {
			return accomodations, err
		}
		accomodations = append(accomodations, elem)
	}
	return accomodations, nil
}

func (accomodationRepository *AccomodationRepository) getCollection(key string) *mongo.Collection {
	return accomodationRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
