package accomodation

import (
	"accomodation_service/accomodation/dtos"
	. "accomodation_service/accomodation/model"
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
		err = cur.Decode(&elem)
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

func (accomodationRepository *AccomodationRepository) Create(accomodation Accomodation) (*Accomodation, error) {
	collection := accomodationRepository.getCollection("accomodations")
	res, err := collection.InsertOne(context.TODO(), accomodation)
	if err != nil {
		return nil, err
	}
	accomodation.Id = res.InsertedID.(primitive.ObjectID)
	return &accomodation, nil
}

func (accommodationRepository *AccomodationRepository) FindById(id primitive.ObjectID) (Accomodation, error) {
	collection := accommodationRepository.getCollection("accomodations")
	var accommodation Accomodation
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&accommodation)
	if err != nil {
		return Accomodation{}, err
	}
	return accommodation, nil
}

func (accomodationRepository *AccomodationRepository) SearchAndFilter(params dtos.SearchDto) ([]Accomodation, error) {
	collection := accomodationRepository.getCollection("accomodations")
	opts := options.Find()
	opts.SetLimit(int64(params.Limit))
	opts.SetSkip(int64((params.Page - 1) * params.Limit))
	fmt.Println(params.Filters)
	pipeline := createSearchAndFilterPipeline(params)
	pipeline = append(pipeline, bson.M{"$skip": (params.Page - 1) * params.Limit})
	pipeline = append(pipeline, bson.M{"$limit": params.Limit})
	fmt.Println(pipeline)

	cur, err := collection.Aggregate(context.TODO(), pipeline)
	var result = make([]Accomodation, 0)

	if err != nil {
		return result, err
	}

	for cur.Next(context.TODO()) {
		var acc Accomodation
		err := cur.Decode(&acc)
		if err != nil {
			return result, err
		}
		result = append(result, acc)
	}

	return result, nil
}

func (accommodationRepository *AccomodationRepository) CountTotalForSearchAndFilter(params dtos.SearchDto) int32 {
	collection := accommodationRepository.getCollection("accomodations")
	pipeline := createSearchAndFilterPipeline(params)
	pipeline = append(pipeline, bson.M{"$count": "total_count"})

	cursor, _ := collection.Aggregate(context.TODO(), pipeline)

	var result []bson.M
	if err := cursor.All(context.TODO(), &result); err != nil {
		log.Fatal(err)
	}

	// Extract the total count
	totalCount := result[0]["total_count"].(int32)
	return totalCount
}

func (accomodationRepository *AccomodationRepository) getCollection(key string) *mongo.Collection {
	return accomodationRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}

func (accomodationRepository *AccomodationRepository) UpdatePricing(accomodation Accomodation) error {
	collection := accomodationRepository.getCollection("accomodations")
	filter := bson.M{"_id": accomodation.Id}
	update := bson.M{"$set": bson.M{"pricing": accomodation.Pricing}}

	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func createSearchAndFilterPipeline(params dtos.SearchDto) bson.A {
	searchFilters := bson.M{
		"$match": bson.M{
			"city": bson.M{"$regex": "(?i).*" + params.City + ".*"},
		},
	}

	if params.Guests != -1 {
		searchFilters["$match"].(bson.M)["min_guests"] = bson.M{"$lte": params.Guests}
		searchFilters["$match"].(bson.M)["max_guests"] = bson.M{"$gte": params.Guests}
	}

	if len(params.IncludingIds) > 0 {
		searchFilters["$match._id"] = bson.M{"$in": params.IncludingIds}
	}

	filters := bson.M{
		"$match": bson.M{
			"pricing": bson.M{
				"$elemMatch": bson.M{
					"price": bson.M{
						"$gte": params.Price.From,
						"$lte": params.Price.To,
					},
				},
			},
		},
	}

	optional := bson.A{}
	for _, f := range params.Filters {
		optionalFilter := bson.M{}
		optionalFilter[f] = true
		optional = append(optional, optionalFilter)
	}

	pipeline := bson.A{
		searchFilters,
		filters,
	}
	if len(optional) > 0 {
		pipeline = append(pipeline, bson.M{"$match": bson.M{"$or": optional}})
	}
	return pipeline
}
