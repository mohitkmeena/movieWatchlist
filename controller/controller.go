package controller

import (
	"encoding/json"
	"mongoapi/model"
	"mongoapi/service"
	"net/http"

	"github.com/gorilla/mux"
)


func GetAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	allmovies:=service.GetAll()
	json.NewEncoder(w).Encode(allmovies)

}
func CreateMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")

	var movie model.Netflix
	json.NewDecoder(r.Body).Decode(&movie)

	id:=service.InsertOneMovie(&movie)
	json.NewEncoder(w).Encode(id)
}
func MarkAsWathced(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","POST")
	params:=mux.Vars(r)
	service.UpdateOneMovie(params["id"])
	json.NewEncoder(w).Encode(params["id"])

}
func GetOneMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","GET")
	params:=mux.Vars(r)
	movie:=service.GetOne(params["id"])
	json.NewEncoder(w).Encode(movie)
}
func DeleteOneMovie(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods",http.MethodDelete)
	params:=mux.Vars(r)
	service.DeleteOne(params["id"])
	json.NewEncoder(w).Encode(params["id"])
}
func DeleteAllMovies(w http.ResponseWriter,r *http.Request){
	w.Header().Set("Content-type","application/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods","DELETE")
	
	service.DeleteAll()
	json.NewEncoder(w).Encode("all movies deleted")
}
func Home(w http.ResponseWriter,r *http.Request){
	w.Write([]byte("<h1>keep going</h1>"))
}
