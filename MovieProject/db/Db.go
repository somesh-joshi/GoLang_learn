package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://127.0.0.1:27017"

const dbName = "netflix"
const colName1 = "watchlist"
const colName2 = "actors"
const colName3 = "directors"

var Collection_watchlist *mongo.Collection
var Collection_actors *mongo.Collection
var Collection_directors *mongo.Collection

// MOST IMPORTANT

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")
	Collection_watchlist = client.Database(dbName).Collection(colName1)
	Collection_actors = client.Database(dbName).Collection(colName2)
	Collection_directors = client.Database(dbName).Collection(colName3)

}
