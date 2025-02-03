package router

import (
	"log"
	"mongoapi/controller"
	"net/http"

	"github.com/gorilla/mux"
)
func Routers(){
	r:=mux.NewRouter()
    r.HandleFunc("/", controller.Home).Methods(http.MethodGet)
	r.HandleFunc("/movies",controller.GetAllMovies).Methods(http.MethodGet)
	r.HandleFunc("/movies/{id}",controller.GetOneMovie).Methods(http.MethodGet)
	r.HandleFunc("/movies",controller.CreateMovie).Methods(http.MethodPost)
	r.HandleFunc("/movies",controller.DeleteAllMovies).Methods(http.MethodDelete)
	r.HandleFunc("/movies/{id}",controller.DeleteOneMovie).Methods(http.MethodDelete)
	r.HandleFunc("/movies/{id}",controller.MarkAsWathced).Methods(http.MethodPut)
	log.Fatal( http.ListenAndServe(":8080",r))
	
	
	

}