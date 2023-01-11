package routers

import (
	"github.com/gorilla/mux"
	"github.com/somesh-joshi/MovieProject/controllers"
)

func Router() *mux.Router {
	routers := mux.NewRouter()
	routers.HandleFunc("/movies", controllers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/movies", controllers.CreateMovie).Methods("POST")
	return routers
}
