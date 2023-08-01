package database

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client       *mongo.Client
	Database     *mongo.Database
	DatabaseName string
}

func NewDatabaseConnection(mongoURI string, dbName string) (*MongoInstance, error) {
	client, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))

	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	err = client.Connect(ctx)
	database := client.Database(dbName)

	if err != nil {
		return nil, err
	}

	mongoInstance := &MongoInstance{
		Client:       client,
		Database:     database,
		DatabaseName: dbName,
	}

	return mongoInstance, nil
}
