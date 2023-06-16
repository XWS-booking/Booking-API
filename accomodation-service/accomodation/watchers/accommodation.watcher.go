package watchers

import (
	"accomodation_service/accomodation"
	"accomodation_service/common/messaging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type AccommodationEventWatcher struct {
	DB                      *mongo.Client
	Publisher               messaging.Publisher
	AccommodationRepository accomodation.IAccomodationRepository
}

type OperationType int32

const (
	CREATE OperationType = 0
	UPDATE OperationType = 1
	DELETE OperationType = 2
)

type AccommodationObj struct {
	Id    string `json:"id"`
	Title string `json:"title"`
}

type AccommodationMessage struct {
	Type          OperationType
	Accommodation AccommodationObj
}

type DbEvent struct {
	DocumentKey   documentKey `bson:"documentKey"`
	OperationType string      `bson:"operationType"`
}
type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}

func (watcher *AccommodationEventWatcher) StartWatching(ctx context.Context) error {
	collection := watcher.getCollection("accomodations")
	pipeline := mongo.Pipeline{}
	changeStream, err := collection.Watch(ctx, pipeline)
	if err != nil {
		log.Fatal("Error when watching accommodation change stream!")
		changeStream.Close(context.TODO())
		return err
	}

	defer changeStream.Close(ctx)

	for changeStream.Next(ctx) {
		var obj DbEvent
		err := changeStream.Decode(&obj)
		if err != nil {
			log.Fatal(err.Error())
			continue
		}
		watcher.PublishMessage(obj)
	}
	return nil
}

func (watcher *AccommodationEventWatcher) PublishMessage(event DbEvent) {
	accommodation := AccommodationObj{Id: event.DocumentKey.ID.Hex()}
	if event.OperationType == "delete" {
		message := AccommodationMessage{Type: DELETE, Accommodation: accommodation}
		err := watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if event.OperationType == "insert" {
		existing, err := watcher.AccommodationRepository.FindById(event.DocumentKey.ID)
		if err != nil {
			log.Fatal(err)
			return
		}
		accommodation.Title = existing.Name
		message := AccommodationMessage{Type: CREATE, Accommodation: accommodation}
		fmt.Println(message)
		err = watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
}

func (watcher *AccommodationEventWatcher) getCollection(key string) *mongo.Collection {
	return watcher.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
