package routers

import (
	"github.com/gorilla/mux"
	actorscontrollers "github.com/somesh-joshi/MovieProject/controllers/actors"
	moviescontrollers "github.com/somesh-joshi/MovieProject/controllers/movies"
	directorscontrollers "github.com/somesh-joshi/MovieProject/controllers/directors"
)

func Router() *mux.Router {
	routers := mux.NewRouter()
	routers.HandleFunc("/movies", moviescontrollers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/movies", moviescontrollers.CreateMovie).Methods("POST")
	routers.HandleFunc("/actors", actorscontrollers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/actors", actorscontrollers.CreateMovie).Methods("POST")
	routers.HandleFunc("/directors", directorscontrollers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/directors", directorscontrollers.CreateMovie).Methods("POST")
	return routers
}
