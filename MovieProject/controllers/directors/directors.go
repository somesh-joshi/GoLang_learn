package directorscontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/directors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = db.Collection_directors

func insertOneMovie(director directors.Director) {
	fmt.Println(director)
	inserted, err := collection.InsertOne(context.Background(), director)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}

func getAllMovies() []primitive.M {
	cur, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var directors []primitive.M

	for cur.Next(context.Background()) {
		var director bson.M
		err := cur.Decode(&director)
		if err != nil {
			log.Fatal(err)
		}
		directors = append(directors, director)
	}

	defer cur.Close(context.Background())
	return directors
}

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var director directors.Director
	_ = json.NewDecoder(r.Body).Decode(&director)
	insertOneMovie(director)
	json.NewEncoder(w).Encode(director)

}
