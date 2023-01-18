package moviescontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
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

func findById(id string) (d primitive.M) {
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(objID)
	filter := bson.M{"_id": objID}
	var data movies.Movie
	err := collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	d = bson.M{"_id": data.ID, "movie": data.Movie, "watched": data.Watched, "actors_id": data.Actors_id, "director": data.Director, "rating": data.Rating}
	return d
}

func FindById(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	id := params["id"]
	actor := findById(id)
	json.NewEncoder(w).Encode(actor)
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
	_ = json.NewDecoder(r.Body).Decode(&movie)
	inserted := insertOneMovie(movie)
	json.NewEncoder(w).Encode(inserted.InsertedID)

}
