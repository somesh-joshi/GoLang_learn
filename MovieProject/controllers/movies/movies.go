package moviescontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/movies"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = db.Db.Collection("watchlist")

func insertOneMovie(movie movies.Movie) (id *mongo.InsertOneResult) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
	return inserted
}

func getAllMovies(filter primitive.D, option *options.FindOptions) []primitive.M {
	cur, err := collection.Find(context.Background(), filter, option)
	if err != nil {
		log.Fatal(err)
	}

	var actors []primitive.M

	for cur.Next(context.Background()) {
		var actor bson.M
		err := cur.Decode(&actor)
		if err != nil {
			log.Fatal(err)
		}
		actors = append(actors, actor)
	}

	defer cur.Close(context.Background())
	return actors
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
	params := r.URL.Query()
	var option *options.FindOptions
	filter := bson.D{}
	list := bson.D{}
	if params["movie"] != nil {
		filter = append(filter, bson.E{Key: "movie", Value: params["movie"][0]})
	}
	if params["rating"] != nil {
		rating, _ := strconv.Atoi(params["rating"][0])
		filter = append(filter, bson.E{Key: "rating", Value: rating})
	}
	if params["director"] != nil {
		data, _ := primitive.ObjectIDFromHex(params["director"][0])
		filter = append(filter, bson.E{Key: "director", Value: data})
	}
	if params["actors_id"] != nil {
		data, _ := primitive.ObjectIDFromHex(params["actors_id"][0])
		filter = append(filter, bson.E{Key: "actors_id", Value: data})
	}
	if params["sort"] != nil {
		list = append(list, bson.E{Key: params["sort"][0], Value: 1})
	}
	option = options.Find().SetSort(list)
	allMovies := getAllMovies(filter, option)
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
