package controller

import (
	"github.com/somesh-joshi/MovieProject/Model/actors"
	"github.com/somesh-joshi/MovieProject/Model/directors"
	"github.com/somesh-joshi/MovieProject/Model/movies"
	"github.com/somesh-joshi/MovieProject/db"
)

func insertOneMovie(movie movie.Movie) {
	inserted, err := collection.InsertOne(context.Background(), movie)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 movie in db with id: ", inserted.InsertedID)
}
