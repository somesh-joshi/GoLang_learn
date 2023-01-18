package routers

import (
	"github.com/gorilla/mux"
	actorscontrollers "github.com/somesh-joshi/MovieProject/controllers/actors"
	directorscontrollers "github.com/somesh-joshi/MovieProject/controllers/directors"
	moviescontrollers "github.com/somesh-joshi/MovieProject/controllers/movies"
)

func Router() *mux.Router {
	routers := mux.NewRouter()
	routers.HandleFunc("/movies", moviescontrollers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/movies", moviescontrollers.CreateMovie).Methods("POST")
	routers.HandleFunc("/movies/{id}", moviescontrollers.FindById).Methods("GET")
	routers.HandleFunc("/actors", actorscontrollers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/actors", actorscontrollers.CreateMovie).Methods("POST")
	routers.HandleFunc("/actors/id/{id}", actorscontrollers.FindById).Methods("GET")
	routers.HandleFunc("/actors/dob/{dob}", actorscontrollers.FindByDoB).Methods("GET")
	routers.HandleFunc("/directors", directorscontrollers.GetMyAllMovies).Methods("GET")
	routers.HandleFunc("/directors", directorscontrollers.CreateMovie).Methods("POST")
	routers.HandleFunc("/directors/{id}", directorscontrollers.FindById).Methods("GET")
	routers.HandleFunc("/directors/{dob}", directorscontrollers.FindByDoB).Methods("GET")
	return routers
}
