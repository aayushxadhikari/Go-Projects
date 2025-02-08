package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct{
	ID string `json:"id"`
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"`
}

type Director struct{
	Firstname string `json:"firstname"`
	Lastname string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content=Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content=Type", "application/json")
	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
		movies = append(movies[:index],movies[index+1:]...)
		w.WriteHeader(http.StatusNoContent)
		return
	}
}
http.Error(w, "Movie not found", http.StatusNotFound)
}

func getMovieById(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type","application/json")
	params := mux.Vars(r)
	for _, movie := range movies{
		if movie.ID == params["id"]{
		json.NewEncoder(w).Encode(movie)
		return
	}
}

	http.Error(w, "Movie not found", http.StatusNotFound)
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var movie Movie
	_=json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(1000000))
	movies = append(movies, movie)
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	
	for index, movie := range movies{
		if movie.ID == params["id"]{
			// Remove old movie
			movies = append(movies[:index], movies[index+1:]...)

			//Create new movie with updated data
			var updatedMovie Movie
			_ = json.NewDecoder(r.Body).Decode(&updatedMovie)
			updatedMovie.ID = params["id"]
			movies = append(movies, updatedMovie)

			json.NewEncoder(w).Encode(updatedMovie)
			return

		}

	}
}


func main(){
	r := mux.NewRouter()
	movies = append(movies, Movie{ID: "1", Isbn:"438227", Title:"Green Book", Director: &Director{Firstname: "Peter", Lastname: "Farrelly"}})
	movies = append(movies, Movie{ID: "2", Isbn:"438228", Title:"Seven", Director: &Director{Firstname: "David", Lastname: "Fincher"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovieById).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))
}