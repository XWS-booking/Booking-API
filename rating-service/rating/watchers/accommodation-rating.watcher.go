package watchers

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	"rating_service/messaging"
	"rating_service/rating"
	"time"
)

type AccommodationRatingEventWatcher struct {
	DB               *mongo.Client
	Publisher        messaging.Publisher
	RatingRepository rating.IRatingRepository
}

type OperationType int32

const (
	CREATE OperationType = 0
	UPDATE OperationType = 1
	DELETE OperationType = 2
)

type RatingObj struct {
	Id              string    `json:"id"`
	UserId          string    `json:"userId"`
	AccommodationId string    `json:"accommodationId"`
	Value           int32     `json:"value"`
	CreatedAt       time.Time `json:"createdAt"`
}

type RatingMessage struct {
	Type   OperationType
	Rating RatingObj
}

type DbEvent struct {
	DocumentKey   documentKey `bson:"documentKey"`
	OperationType string      `bson:"operationType"`
}
type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}

func (watcher *AccommodationRatingEventWatcher) StartWatching(ctx context.Context) error {
	collection := watcher.getCollection("accommodation_ratings")
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

func (watcher *AccommodationRatingEventWatcher) PublishMessage(event DbEvent) {
	fmt.Println(event)
	rating := RatingObj{Id: event.DocumentKey.ID.Hex()}

	if event.OperationType == "delete" {
		message := RatingMessage{Type: DELETE, Rating: rating}
		err := watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if event.OperationType == "insert" {
		existing, err := watcher.RatingRepository.FindAccommodationRatingById(event.DocumentKey.ID)
		if err != nil {
			log.Fatal(err)
			return
		}
		rating.CreatedAt = existing.Time
		rating.Value = existing.Rating
		rating.UserId = existing.GuestId.Hex()
		rating.AccommodationId = existing.AccommodationId.Hex()
		message := RatingMessage{Type: CREATE, Rating: rating}
		fmt.Println(message)
		err = watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if event.OperationType == "update" {
		existing, err := watcher.RatingRepository.FindAccommodationRatingById(event.DocumentKey.ID)
		if err != nil {
			log.Fatal(err)
			return
		}
		rating.CreatedAt = existing.Time
		rating.Value = existing.Rating
		rating.UserId = existing.GuestId.Hex()
		rating.AccommodationId = existing.AccommodationId.Hex()
		message := RatingMessage{Type: UPDATE, Rating: rating}
		fmt.Println(message)
		err = watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
}

func (watcher *AccommodationRatingEventWatcher) getCollection(key string) *mongo.Collection {
	return watcher.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
