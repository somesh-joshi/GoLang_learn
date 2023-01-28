package person

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/somesh-joshi/MovieProject/db"
	"github.com/somesh-joshi/MovieProject/models/person"
	"github.com/somesh-joshi/MovieProject/validator"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection = db.Db.Collection("person")

func insertOneMovie(person person.Person) (id *mongo.InsertOneResult) {
	inserted, err := collection.InsertOne(context.Background(), person)

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

func findById(filter primitive.D) (d primitive.M) {
	var data person.Person
	err := collection.FindOne(context.Background(), filter).Decode(&data)
	if err != nil {
		log.Fatal(err)
	}
	d = bson.M{"id": data.ID, "name": data.Name, "age": data.Age, "dob": data.DoB}
	return d
}

// exportable func

func FindByIdActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	id := params["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{bson.E{Key: "title", Value: "actor"}, bson.E{Key: "_id", Value: objID}}
	actor := findById(filter)
	json.NewEncoder(w).Encode(actor)
}

func FindByIdDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	params := mux.Vars(r)
	id := params["id"]
	objID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.D{bson.E{Key: "title", Value: "director"}, bson.E{Key: "_id", Value: objID}}
	actor := findById(filter)
	json.NewEncoder(w).Encode(actor)
}

func GetAllActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-url")
	params := r.URL.Query()
	var option *options.FindOptions
	filter := bson.D{bson.E{Key: "title", Value: "actor"}}
	list := bson.D{}
	if params["dob"] != nil {
		filter = append(filter, bson.E{Key: "dob", Value: params["dob"][0]})
	}
	if params["age"] != nil {
		age, _ := strconv.Atoi(params["age"][0])
		filter = append(filter, bson.E{Key: "age", Value: age})
	}
	if params["name"] != nil {
		filter = append(filter, bson.E{Key: "name", Value: params["name"][0]})
	}
	if params["sort"] != nil {
		list = append(list, bson.E{Key: params["sort"][0], Value: 1})
	}
	option = options.Find().SetSort(list).SetProjection(bson.D{{Key: "title", Value: 0}})
	allMovies := getAllMovies(filter, option)
	json.NewEncoder(w).Encode(allMovies)
}

func GetAllDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-url")
	params := r.URL.Query()
	var option *options.FindOptions
	filter := bson.D{bson.E{Key: "title", Value: "director"}}
	list := bson.D{}
	if params["dob"] != nil {
		filter = append(filter, bson.E{Key: "dob", Value: params["dob"][0]})
	}
	if params["age"] != nil {
		age, _ := strconv.Atoi(params["age"][0])
		filter = append(filter, bson.E{Key: "age", Value: age})
	}
	if params["name"] != nil {
		filter = append(filter, bson.E{Key: "name", Value: params["name"][0]})
	}
	if params["sort"] != nil {
		list = append(list, bson.E{Key: params["sort"][0], Value: 1})
	}
	option = options.Find().SetSort(list).SetProjection(bson.D{{Key: "title", Value: 0}})
	allMovies := getAllMovies(filter, option)
	json.NewEncoder(w).Encode(allMovies)
}

func CreateActor(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var actor person.Person
	_ = json.NewDecoder(r.Body).Decode(&actor)
	actor.Title = "actor"
	err := validator.Validator(actor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	} else {
		inserted := insertOneMovie(actor)
		json.NewEncoder(w).Encode(inserted.InsertedID)
	}
}

func CreateDirector(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var actor person.Person
	_ = json.NewDecoder(r.Body).Decode(&actor)
	actor.Title = "director"
	err := validator.Validator(actor)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(err.Error())
	} else {
		inserted := insertOneMovie(actor)
		json.NewEncoder(w).Encode(inserted.InsertedID)
	}
}
