package util

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Client       *mongo.Client
	DBNAME       = "authentication"
	DBCollection = "users"
)

func ConnectDB() (*mongo.Client, error) {
	c, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	client, err := mongo.Connect(c, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return client, err
	}
	Client = client

	return Client, nil
}
