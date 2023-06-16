package database

import (
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"os"
)

func InitDB() (neo4j.DriverWithContext, error) {
	uri := os.Getenv("NEO4J_URI")
	username := os.Getenv("NEO4J_USERNAME")
	password := os.Getenv("NEO4J_PASSWORD")
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(username, password, ""))
	if err != nil {
		return nil, err
	}
	return driver, nil
}
