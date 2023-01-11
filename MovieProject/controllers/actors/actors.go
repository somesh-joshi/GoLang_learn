package actorscontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/actors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var collection = db.Collection_actors

func insertOneMovie(actor actors.Actor) {
	fmt.Println(actor)
	inserted, err := collection.InsertOne(context.Background(), actor)

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

func GetMyAllMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	allMovies := getAllMovies()
	json.NewEncoder(w).Encode(allMovies)
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var actor actors.Actor
	_ = json.NewDecoder(r.Body).Decode(&actor)
	insertOneMovie(actor)
	json.NewEncoder(w).Encode(actor)

}
