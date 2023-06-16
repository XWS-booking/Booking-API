package accommodation

import (
	"context"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"recommendation_service/recommendation/model"
)

type AccommodationRepository struct {
	Db neo4j.DriverWithContext
}

func (accommodationRepository *AccommodationRepository) Create(accommodation model.Accommodation) error {
	ctx := context.TODO()
	query := "CREATE (u:Accommodation{_id:$id, title:$title})"
	params := map[string]interface{}{"id": accommodation.Id, "title": accommodation.Title}

	session := accommodationRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})
	return err
}

func (accommodationRepository *AccommodationRepository) Delete(accommodation model.Accommodation) error {
	ctx := context.TODO()
	query := "MATCH (u:Accommodation{_id:$id})-[r]-() delete r, u"
	params := map[string]interface{}{"id": accommodation.Id}

	session := accommodationRepository.Db.NewSession(ctx, neo4j.SessionConfig{})
	_, err := session.ExecuteWrite(ctx, func(transaction neo4j.ManagedTransaction) (any, error) {
		return transaction.Run(ctx, query, params)
	})
	return err
}
