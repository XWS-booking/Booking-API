package watchers

import (
	"auth_service/auth"
	"auth_service/common/messaging"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
)

type UserEventWatcher struct {
	DB             *mongo.Client
	Publisher      messaging.Publisher
	UserRepository auth.IUserRepository
}

type OperationType int32

const (
	CREATE OperationType = 0
	UPDATE OperationType = 1
	DELETE OperationType = 2
)

type UserObj struct {
	Id string `json:"id"`
}

type UserMessage struct {
	Type OperationType
	User UserObj
}

type DbEvent struct {
	DocumentKey   documentKey `bson:"documentKey"`
	OperationType string      `bson:"operationType"`
}
type documentKey struct {
	ID primitive.ObjectID `bson:"_id"`
}

func (watcher *UserEventWatcher) StartWatching(ctx context.Context) error {
	collection := watcher.getCollection("users")
	pipeline := mongo.Pipeline{}
	changeStream, err := collection.Watch(ctx, pipeline)
	if err != nil {
		log.Fatal("Error when watching user change stream!")
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

		fmt.Println("Change type is ", obj.OperationType, obj.DocumentKey)
	}
	return nil
}

func (watcher *UserEventWatcher) PublishMessage(event DbEvent) {
	user := UserObj{Id: event.DocumentKey.ID.Hex()}
	if event.OperationType == "delete" {
		message := UserMessage{Type: DELETE, User: user}
		fmt.Println("Publishing delete ", message)

		err := watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}

	if event.OperationType == "insert" {
		message := UserMessage{Type: CREATE, User: user}
		fmt.Println("Publishing inser ", message)

		err := watcher.Publisher.Publish(&message)
		if err != nil {
			log.Fatal(err.Error())
		}
		return
	}
}

func (watcher *UserEventWatcher) getCollection(key string) *mongo.Collection {
	return watcher.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
