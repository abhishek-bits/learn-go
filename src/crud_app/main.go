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

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"string"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	// Set the Request Headers
	w.Header().Set("Content-Type", "appication/json")

	// Encode the response into JSON
	json.NewEncoder(w).Encode(movies)
}

func deleteMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// Fetch the params of the API
	params := mux.Vars(r)

	for index, item := range movies {
		if item.ID == params["id"] {
			// We can use append() and basic slice operation
			// to get rid of the movie to be deleted.
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}

	// Return the remaining slice of movies
	json.NewEncoder(w).Encode(movies)
}

func getMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)

	// We don't need index this time.
	for _, item := range movies {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func createMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	var movie Movie

	// We have the request in JSON format
	// We simply decode it and populate
	// our movie variable.
	_ = json.NewDecoder(r.Body).Decode(&movie)

	movie.ID = strconv.Itoa(rand.Intn(10000000))

	// append this movie into movies list
	movies = append(movies, movie)

	// return the newly created movie
	json.NewEncoder(w).Encode(movie)
}

func updateMovie(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")

	// ID is passed in the params
	params := mux.Vars(r)

	var movie Movie

	// We'll use JSON decoder to decode the
	// request body into a movie type.
	_ = json.NewDecoder(r.Body).Decode(&movie)

	for index, item := range movies {
		if item.ID == params["id"] {
			// Remove the old movie from the slice
			movies = append(movies[:index], movies[index+1:]...)
			// Set the ID for the updated movie
			movie.ID = params["id"]
			// Append the updated movie to the slice
			movies = append(movies, movie)
			break
		}
	}

	json.NewEncoder(w).Encode(movies)
}

func main() {

	r := mux.NewRouter()

	movies = append(
		movies,
		Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}},
		Movie{ID: "2", Isbn: "454556", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}},
	)

	r.HandleFunc("/movies", getMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")

	// Create a web server
	log.Fatal(http.ListenAndServe(":8000", r))

}
