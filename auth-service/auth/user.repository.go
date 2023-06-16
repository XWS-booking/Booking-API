package auth

import (
	. "auth_service/auth/model"
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (userRepository *UserRepository) Create(user User) (primitive.ObjectID, error) {
	collection := userRepository.getCollection("users")
	res, err := collection.InsertOne(context.TODO(), user)
	if err != nil {
		userRepository.Logger.Println(err)
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (userRepository *UserRepository) FindById(id primitive.ObjectID) (User, error) {
	collection := userRepository.getCollection("users")
	var user User
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		userRepository.Logger.Println(err)
		return User{}, err
	}
	return user, nil
}

func (userRepository *UserRepository) FindByEmail(email string) (User, error) {
	collection := userRepository.getCollection("users")
	var user User
	filter := bson.M{"email": email}
	err := collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		userRepository.Logger.Println(err)
		return user, err
	}
	return user, nil
}

func (userRepository *UserRepository) Delete(id primitive.ObjectID) error {
	collection := userRepository.getCollection("users")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (userRepository *UserRepository) getCollection(key string) *mongo.Collection {
	return userRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}

func (userRepository *UserRepository) UpdatePersonalInfo(user User) (User, error) {
	collection := userRepository.getCollection("users")
	filter := bson.M{"_id": user.Id}
	update := bson.M{"$set": user}
	_, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		userRepository.Logger.Println(err)
		return user, err
	}
	return user, nil
}

func (userRepository *UserRepository) GetFeaturedHosts() ([]User, error) {
	collection := userRepository.getCollection("users")
	filter := bson.M{"distinguished": true}

	curr, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		return []User{}, nil
	}
	result := make([]User, 0)
	for curr.Next(context.TODO()) {
		var user User
		err := curr.Decode(&user)
		if err != nil {
			log.Fatal(err)
			return []User{}, err
		}
		result = append(result, user)
	}
	return result, nil
}
