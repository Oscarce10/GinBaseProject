package config

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"os"
)

type Database struct {
	Client *mongo.Client
	Err    error
}

func (db *Database) GetClient() *mongo.Client {
	db.Client, db.Err = mongo.Connect(
		context.Background(),
		options.Client().ApplyURI(os.Getenv("MONGODB_URI")),
	)
	if db.Err != nil {
		log.Fatal("Error connecting to MongoDB: ", db.Err.Error())
	}
	return db.Client
}

func (db *Database) GetCollection(collection string) *mongo.Collection {
	db.Err = db.Client.Ping(context.Background(), nil)
	if db.Err != nil {
		log.Fatal("Error connecting to MongoDB: ", db.Err.Error())
	}
	return db.Client.Database(os.Getenv("MONGODB_DATABASE")).Collection(collection)
}

func (db *Database) CloseClient() {
	if db.Client != nil {
		db.Err = db.Client.Disconnect(context.TODO())
		if db.Err != nil {
			log.Fatal("Error disconnecting from MongoDB: ", db.Err.Error())
		}
	}
	return
}
