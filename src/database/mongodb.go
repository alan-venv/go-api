package database

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongodb *mongo.Database

func Start() {
	//! Set client options
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_CONNSTRING"))

	//! Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal("[MongoDB] Cannot connect to database!")
	}

	//! Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("[MongoDB] Fail to ping on primary database!")
	}

	mongodb = client.Database("goapi")
}

func Get() *mongo.Database {
	return mongodb
}
