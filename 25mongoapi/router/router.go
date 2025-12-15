package router

import (
	"github.com/gorilla/mux"
	"github.com/khubaibkhalil/mongoapi/controller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/movies", controller.GetAllMovies).Methods("GET")

	r.HandleFunc("/api/movie", controller.CreateMovie).Methods("POST")

	r.HandleFunc("/api/movie/{id}", controller.MarkAsWatched).Methods("PUT")

	r.HandleFunc("/api/movie/{id}", controller.DeleteAMovie).Methods("DELETE")

	r.HandleFunc("/api/movies/deleteall", controller.DeleteAllMovies).Methods("DELETE")

	return r
}
