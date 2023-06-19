package rating

import (
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"recommendation_service/recommendation/model"
)

type RatingRepository struct {
	Db neo4j.DriverWithContext
}

func (ratingRepository *RatingRepository) Create(rating model.Rating) error {
	ctx := context.TODO()
	query := "MATCH (u:User{_id:$userId}), (a:Accommodation{_id:$accId}) " +
		"CREATE (u)-[r:RATES{_id: $id,value:$value, createdAt:$createdAt}]->(a)"
	params := map[string]interface{}{
		"id":        rating.Id,
		"userId":    rating.UserId,
		"accId":     rating.AccommodationId,
		"value":     rating.Value,
		"createdAt": rating.CreatedAt,
	}
	fmt.Println(params)
	session := ratingRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})
	if err != nil {
		fmt.Println("Something wrong happened")
	}
	return err
}

func (ratingRepository *RatingRepository) Delete(rating model.Rating) error {
	ctx := context.TODO()
	query := "MATCH ()-[r:RATES{_id:$id}]->() delete r"
	params := map[string]interface{}{
		"id": rating.Id,
	}
	session := ratingRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})
	return err
}

func (ratingRepository *RatingRepository) Update(rating model.Rating) error {
	ctx := context.TODO()
	query := "MATCH (u:User{_id:$userId})-[r:RATES{_id:$id}]->(a:Accommodation{_id:$accId}) " +
		"SET r.value=$value, r.createdAt=$createdAt"
	params := map[string]interface{}{
		"id":        rating.Id,
		"userId":    rating.UserId,
		"accId":     rating.AccommodationId,
		"value":     rating.Value,
		"createdAt": rating.CreatedAt,
	}
	session := ratingRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})

	log.Fatal(err)
	return err
}
