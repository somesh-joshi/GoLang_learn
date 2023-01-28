package routers

import (
	"net/http"

	"github.com/gorilla/mux"

	auth "github.com/somesh-joshi/MovieProject/auth"
	moviescontrollers "github.com/somesh-joshi/MovieProject/controllers/movies"
	personcontrollers "github.com/somesh-joshi/MovieProject/controllers/person"
)

func Router() *mux.Router {
	routers := mux.NewRouter()
	routers.HandleFunc("/signup", auth.Logup).Methods("POST")
	routers.HandleFunc("/signin", auth.Login).Methods("POST")
	routers.Handle("/actors", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.GetAllActor))).Methods("GET")
	routers.Handle("/movies", auth.AuthMiddleware(http.HandlerFunc(moviescontrollers.GetMyAllMovies))).Methods("GET")
	routers.Handle("/movies", auth.AuthMiddleware(http.HandlerFunc(moviescontrollers.CreateMovie))).Methods("POST")
	routers.Handle("/movies/{id}", auth.AuthMiddleware(http.HandlerFunc(moviescontrollers.FindById))).Methods("GET")
	routers.Handle("/actors", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.GetAllActor))).Methods("GET")
	routers.Handle("/actors", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.CreateActor))).Methods("POST")
	routers.Handle("/actors/{id}", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.FindByIdActor))).Methods("GET")
	routers.Handle("/directors", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.GetAllDirector))).Methods("GET")
	routers.Handle("/directors", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.CreateDirector))).Methods("POST")
	routers.Handle("/directors/{id}", auth.AuthMiddleware(http.HandlerFunc(personcontrollers.FindByIdDirector))).Methods("GET")
	return routers
}
