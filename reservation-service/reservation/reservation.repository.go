package reservation

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"os"
	. "reservation_service/reservation/model"
	"time"
)

type ReservationRepository struct {
	DB     *mongo.Client
	Logger *log.Logger
}

func (reservationRepository *ReservationRepository) FindById(id primitive.ObjectID) (Reservation, error) {
	collection := reservationRepository.getCollection("reservations")
	var reservation Reservation
	filter := bson.M{"_id": id}
	err := collection.FindOne(context.TODO(), filter).Decode(&reservation)
	if err != nil {
		return Reservation{}, err
	}
	return reservation, nil
}

func (reservationRepository *ReservationRepository) Create(reservation Reservation) (primitive.ObjectID, error) {
	collection := reservationRepository.getCollection("reservations")
	res, err := collection.InsertOne(context.TODO(), reservation)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return res.InsertedID.(primitive.ObjectID), nil
}

func (reservationRepository *ReservationRepository) Delete(id primitive.ObjectID) error {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{"_id": id}
	_, err := collection.DeleteOne(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (reservationRepository *ReservationRepository) FindAllReservedAccommodations(startDate time.Time, endDate time.Time) ([]string, error) {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{
		"$and": []bson.M{
			bson.M{"start_date": bson.M{"$lte": endDate}},
			bson.M{"end_date": bson.M{"$gte": startDate}},
		},
		"status": Confirmed,
	}
	field := "accommodation_id"
	distinctValues, err := collection.Distinct(context.TODO(), field, filter)
	if err != nil {
		return []string{}, err
	}

	var accommodationIds []string
	for _, value := range distinctValues {
		id := value.(primitive.ObjectID)
		accommodationIds = append(accommodationIds, id.Hex())
	}
	return accommodationIds, nil
}

func (reservationRepository *ReservationRepository) CheckActiveReservationsForGuest(id primitive.ObjectID) (bool, error) {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{
		"buyer_id": id,
		"end_date": bson.M{
			"$gt": time.Now(),
		},
		"status": bson.M{"$in": bson.A{Pending, Confirmed}},
	}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (reservationRepository *ReservationRepository) CheckActiveReservationsForAccommodation(id primitive.ObjectID) (bool, error) {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{
		"accommodation_id": id,
		"end_date": bson.M{
			"$gt": time.Now(),
		},
		"status": bson.M{"$in": bson.A{Pending, Confirmed}},
	}

	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}

func (reservationRepository *ReservationRepository) DeleteByBuyerId(id primitive.ObjectID) error {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{"buyer_id": id}
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (reservationRepository *ReservationRepository) DeleteByAccommodationId(id primitive.ObjectID) error {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{"accommodation_id": id}
	_, err := collection.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}
	return nil
}

func (reservationRepository *ReservationRepository) UpdateReservation(reservation Reservation) error {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{"_id": reservation.Id}
	_, err := collection.UpdateOne(context.TODO(), filter, bson.M{"$set": reservation})
	if err != nil {
		return err
	}
	return nil
}

func (reservationRepository *ReservationRepository) IsAccommodationAvailable(id primitive.ObjectID, startDate time.Time, endDate time.Time) (bool, error) {
	collection := reservationRepository.getCollection("reservations")
	filter := bson.M{
		"accommodation_id": id,
		"$and": []bson.M{
			bson.M{"start_date": bson.M{"$lte": endDate}},
			bson.M{"end_date": bson.M{"$gte": startDate}}},
		"status": Confirmed}

	count, err := collection.CountDocuments(context.TODO(), filter)

	if err != nil {
		return false, err
	}
	return count == 0, nil
}

func (reservationRepository *ReservationRepository) FindAllByBuyerId(id primitive.ObjectID) ([]Reservation, error) {
	collection := reservationRepository.getCollection("reservations")
	var reservations []Reservation

	filter := bson.M{"buyer_id": id}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return reservations, err
	}

	for cur.Next(context.TODO()) {
		var elem Reservation
		err := cur.Decode(&elem)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, elem)
	}
	return reservations, nil
}

func (reservationRepository *ReservationRepository) FindNumberOfBuyersCancellations(id primitive.ObjectID) (int, error) {
	collection := reservationRepository.getCollection("reservations")
	var reservations []Reservation

	filter := bson.M{"buyer_id": id, "status": Status(3)}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return 0, err
	}

	for cur.Next(context.TODO()) {
		var elem Reservation
		err := cur.Decode(&elem)
		if err != nil {
			return 0, err
		}
		reservations = append(reservations, elem)
	}
	return len(reservations), nil
}

func (reservationRepository *ReservationRepository) FindAllByAccommodationId(id primitive.ObjectID) ([]Reservation, error) {
	collection := reservationRepository.getCollection("reservations")
	var reservations []Reservation

	filter := bson.M{"accommodation_id": id}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return reservations, err
	}

	for cur.Next(context.TODO()) {
		var elem Reservation
		err := cur.Decode(&elem)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, elem)
	}
	return reservations, nil
}

func (reservationRepository *ReservationRepository) FindAllPendingByAccommodationId(id primitive.ObjectID) ([]Reservation, error) {
	collection := reservationRepository.getCollection("reservations")
	var reservations []Reservation

	filter := bson.M{"accommodation_id": id, "status": Status(0)}
	cur, err := collection.Find(context.TODO(), filter)
	if err != nil {
		return reservations, err
	}

	for cur.Next(context.TODO()) {
		var elem Reservation
		err := cur.Decode(&elem)
		if err != nil {
			return reservations, err
		}
		reservations = append(reservations, elem)
	}
	return reservations, nil
}

func (reservationRepository *ReservationRepository) getCollection(key string) *mongo.Collection {
	return reservationRepository.DB.Database(os.Getenv("DATABASE_NAME")).Collection(key)
}
