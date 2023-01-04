package controller

import (
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/somesh-joshi/mongoapi/model"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/mongo/driver/mongocrypt/options"

)

const connectionString = "mongodb://localhost:27017"
const dbName = "movies"
const collectionName = "wmovies"

var client *mongo.Collection

func init() {
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	collection := client.Database(dbName).Collection(collectionName)
}