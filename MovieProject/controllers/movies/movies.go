package moviescontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/movies"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = db.Db.Collection("watchlist")

//var collection_actor = db.Collection_actors
//var collection_director = db.Collection_directors

func insertOneMovie(movie movies.Movie) (id *mongo.InsertOneResult) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
	return inserted
}

//get actor name from actor id

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var movies []primitive.M

	for cur.Next(context.Background()) {
		var movie bson.M
		err := cur.Decode(&movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, movie)
	}

	defer cur.Close(context.Background())
	return movies
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var movie movies.Movie
	inserted := insertOneMovie(movie)
	json.NewEncoder(w).Encode(inserted.InsertedID)

}
