package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID    string `json:"id"`
	Isbn  string `json:"isbn"`
	Title string `json:"title"`
	// This pointer is pointing to the Director struct
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

// Create a slice(array of undefined size) of the Movie struct
var movies []Movie

func main() {
	r := mux.NewRouter()

	// This line could be written using new(Director) istead of &Director if it was started with zero values
	// 	movies = append(movies, Movie{ID: "1", Isbn: "743843", Title: "Test Movie 1", Director: &Director{Firstname: "Director", Lastname: "One"}})
	movies = append(movies, Movie{ID: "1", Isbn: "743843", Title: "Test Movie 1", Director: &Director{Firstname: "Director", Lastname: "One"}})
	// The & is creating a pointer to Director and to dereference this pointer just use *
	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovies).Methods("GET")
	r.HandleFunc("/movies", createMovies).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting the server inport 8000\n")
	// To start the server is very simple using the http library
	log.Fatal(http.ListenAndServe(":8000", r))
}
