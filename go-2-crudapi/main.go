package main

import (
	"encoding/json" // to encode data to json
	"fmt"
	"log"
	"math/rand"
	"net/http" // what creates the webserver
	"strconv"

	"github.com/gorilla/mux" // package to create routes
)

// struct like an object key:value pairs:value

type Movie struct {
	ID string `json:"id"` // this is to encode and decode the jsons to structs and viceversa
	Isbn string `json:"isbn"`
	Title string `json:"title"`
	Director *Director `json:"director"` // * pointer
}

type Director struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	for index, item := range movies {
		if item.ID == params["id"]{
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}


func getMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params:= mux.Vars(r)
	for _, item := range movies {
		if item.ID == params["id"]{
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	
	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	for index, item  := range movies {
		if item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			
			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = item.ID
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID:"1", Isbn: "438227", Title: "Movie one", Director: &Director{FirstName: "John", LastName: "Doe"}})
	movies = append(movies, Movie{ID:"2", Isbn: "438228", Title: "Movie two", Director: &Director{FirstName: "Steve", LastName: "Smith"}})

	// routes
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000",r))



}