package user

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"log"
	"recommendation_service/recommendation/model"
)

type UserRepository struct {
	Db neo4j.DriverWithContext
}

func (userRepository *UserRepository) Create(user model.User) error {
	ctx := context.TODO()
	query := "CREATE (u:User{_id:$id})"
	params := map[string]interface{}{"id": user.Id}

	session := userRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})
	return err
}

func (userRepository *UserRepository) Delete(user model.User) error {
	ctx := context.TODO()
	query := "MATCH (u:User{_id:$id})-[r]-() delete r, u"
	params := map[string]interface{}{"id": user.Id}

	session := userRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})
	return err
}

func (userRepository *UserRepository) GetRecommended(user model.User) ([]model.Accommodation, error) {
	ctx := context.TODO()
	query := "MATCH (p1:User {_id:$id})-[r1:RATING]->(:Accommodation)<-[r2:RATING]-(p2:User) " +
		"WHERE p1 <> p2 and abs(r1.value - r2.value) in [0, 1] " +
		"WITH p2 " +
		"MATCH (p2)-[r:RATING]->(a2:Accommodation) " +
		"WHERE  r.value >= 3 " +
		"WITH DISTINCT a2 " +
		"OPTIONAL MATCH (a2)<-[r:RATING]-() " +
		"WHERE r.value < 3 and r.createdAt < datetime().epochSeconds() - (90 * 24 * 60 * 60) " +
		"WITH a2, count(r) AS lowRatingCount " +
		"MATCH (a2) " +
		"WHERE lowRatingCount = 0 " +
		"WITH a2 " +
		"MATCH (a2)<-[r:RATING]-(:User) " +
		"WITH a2, avg(r.value) as avgRating " +
		"RETURN a2 " +
		"ORDER BY avgRating DESC " +
		"LIMIT 10"
	params := map[string]interface{}{"id": user.Id}

	session := userRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	result, err := session.ExecuteRead(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		result, err := transaction.Run(ctx, query, params)
		results := make([]model.Accommodation, 0)
		if err != nil {
			log.Fatal(err)
			return []model.Accommodation{}, err
		}

		for result.Next(ctx) {
			record := result.Record()
			id, err := record.Get("id")
			if err {
				break
			}
			results = append(results, model.Accommodation{Id: id.(string)})
		}
		return results, nil
	})

	return result.([]model.Accommodation), err

}
