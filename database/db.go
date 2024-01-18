package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func StartDB() error {
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environment variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}
	var err error
	MongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return fmt.Errorf("could not start mongodb client %w", err)
	}
	return nil
}

func CloseDB() error {
	if err := MongoClient.Disconnect(context.TODO()); err != nil {
		panic(err)
	}
	return nil
}

func GetCollection(database string, collection string) *mongo.Collection {
	return MongoClient.Database(database).Collection(collection)
}
