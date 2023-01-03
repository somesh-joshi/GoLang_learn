package controller

import (
	"context"
	"fmt"
	"log"

	"github.com/somesh-joshi/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const connectionString = "mongodb://localhost:27017"
const DBName = "netflix"
const CollectionName = "movies"

var collection *mongo.Collection

func init() {
	// Set client options
	clientOptions := options.Client().ApplyURI(connectionString)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection = client.Database(DBName).Collection(CollectionName)
}

// monogodb helper

// InsertOne insert one document

func InsertOne(movie model.Netflix) (*mongo.InsertOneResult, error) {
	inserted, err := collection.InsertOne(context.Background(), movie)
	if err != nil {
		log.Fatal(err)
	}

	return inserted, err
}

// FindAll find all documents

func FindAll() ([]model.Netflix, error) {
	var movies []model.Netflix
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		log.Fatal(err)
	}

	for cursor.Next(context.Background()) {
		var movie model.Netflix
		cursor.Decode(&movie)
		movies = append(movies, movie)
	}

	return movies, err
}

// FindOne find one document

func FindOne(id string) (model.Netflix, error) {
	var movie model.Netflix
	err := collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&movie)
	if err != nil {
		log.Fatal(err)
	}
	return movie, err
}

// update movie



// delete movie

func DeleteOne(id string) (*mongo.DeleteResult, error) {
	deleted, err := collection.DeleteOne(context.Background(), bson.M{"_id": id})
	if err != nil {
		log.Fatal(err)
	}

	return deleted, err
}
