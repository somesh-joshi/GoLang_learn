package directorscontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/directors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection = db.Db.Collection("directors")

func insertOneMovie(director directors.Director) (id *mongo.InsertOneResult) {
	fmt.Println(director)
	inserted, err := collection.InsertOne(context.Background(), director)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
	return inserted
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
func findById(id string) (d primitive.M) {
	objID, _ := primitive.ObjectIDFromHex(id)
	fmt.Println(objID)
	filter := bson.M{"_id": objID}
	var data directors.Director
	err := collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	d = bson.M{"_id": data.ID, "name": data.Name, "age": data.Age, "dob": data.DoB}
	return d
}

func findBydob(dob string) (d primitive.M) {
	filter := bson.M{"dob": dob}
	var data directors.Director
	err := collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	d = bson.M{"_id": data.ID, "name": data.Name, "age": data.Age, "dob": data.DoB}
	return d
}

func FindByDoB(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	dob := params["dob"]
	fmt.Println(dob)
	actor := findBydob(dob)
	json.NewEncoder(w).Encode(actor)
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

	var director directors.Director
	_ = json.NewDecoder(r.Body).Decode(&director)
	inserted := insertOneMovie(director)
	json.NewEncoder(w).Encode(inserted.InsertedID)

}
