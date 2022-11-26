package main

// Libs
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	// Using the gorilla/mux 3rd party router package instead of the standard library
	// net/http router. Allows you to more easily perform tasks such as parsing path
	// or query params.
	"github.com/gorilla/mux"
)

// Local
import "pkg/tutorial"

const PORT = ":10000"

type Movie struct {
	Id          string `json:"Id"`
	Title       string `json:"Title"`
	Desc        string `json:"Desc"`
	ReleaseYear int    `json:"ReleaseYear"`
}

// NOTE: updating a global variable (MovieList) in order to keep this simple. Not
// doing any checks to ensure race conditions don't happen. This code isn't "thread-safe".

// Global Movies array. Can populate in the `main` function to simulate a db
var MovieList []Movie

// Descriptions
const allAtOnceDesc = "When an interdimensional rupture unravels reality, an unlikely hero must channel her newfound powers to fight bizarre and..."
const troopersDesc = "Five Vermont state troopers, avid pranksters with a knack for screwing up, try to save their jobs and out-do the local police..."

// IMPORTANT:
// Matches the URL path hit with a defined function
func registerHandlers() {
	// Create a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	// Add "/movies" endpoint & map it to the getMovieList ƒn
	myRouter.HandleFunc("/movies", getMovieList)
	// NOTE: Order matters. Must be before the other "/movie" endpoint
	myRouter.HandleFunc("/movie", addMovie).Methods("POST")
	// Order still matters.
	myRouter.HandleFunc("/movie/{id}", deleteMovie).Methods("DELETE")
	myRouter.HandleFunc("/movie/{id}", updateMovie).Methods("PUT")
	// {id} = path variable
	myRouter.HandleFunc("/movie/{id}", getMovie)

	log.Fatal(http.ListenAndServe(PORT, myRouter))
}

// Obvi most important ƒn ✨
func main() {
	// Will execute when you `go run` this file
	fmt.Println("Mux Routers 🦊")

	tutorial.Tutorial()

	var jsonMovies []Movie
	// Get an environment variable in Go.
	if os.Getenv("GOPATH") == "/go" && os.Getenv("HOME") == "/root" {
		jsonMovies = readMoviesFromJSON()
	}
	// Else it'll have zero values

	MovieList = []Movie{
		{Id: "1", Title: "Everything Everywhere All at Once", Desc: allAtOnceDesc, ReleaseYear: 2022},
		{Id: "2", Title: "Super Troopers", Desc: troopersDesc, ReleaseYear: 2001},
		{Id: "3", Title: "3rd title", Desc: "Another movie", ReleaseYear: 1998},
	}

	// Concatenate two SLICES
	// golang's equivalent to javascript's spread operator
	MovieList = append(MovieList, jsonMovies...)

	registerHandlers()
}

// SECTION: Route Handlers

// Handles requests to the root URL.
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the homepage of my first Go repo 🂡")
	fmt.Println("Endpoint hit: homepage")
}

// "R" of CRUD
func getMovieList(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /movies")
	// Encodes the movies into a JSON string
	json.NewEncoder(w).Encode(MovieList)
}

// "R" of CRUD
func getMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit: /movie/{id}")

	routeVariables := mux.Vars(r)
	key := routeVariables["id"]

	// Loop over MovieList
	for _, movie := range MovieList {
		if movie.Id == key {
			// Return the movie encoded as JSON
			json.NewEncoder(w).Encode(movie)
		}
	}
}

// the "C" of CRUD
// ... doesn't take validation into consideration.
func addMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit (POST request received): /movie")

	postBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem reading request body")
		return
	}
	// Unmarshal the request body JSON into a new `Movie` struct.
	var movie Movie
	err = json.Unmarshal(postBody, &movie)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem unmarshalling")
	}

	// Update the global MovieList to include the new movie.
	MovieList = append(MovieList, movie)

	// %+v  -> value in a default format, with field name.
	// use when printing structs
	fmt.Fprintf(w, "%+v", string(postBody))
}

// the "D" of CRUD
func deleteMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit (DELETE request received): /movie/{id}")

	// Extract the ID
	routeVariables := mux.Vars(r)
	deleteId := routeVariables["id"]

	// Loop the movies, remove any entry whose Id property matches deleteId
	for index, movie := range MovieList {
		if movie.Id == deleteId {
			MovieList = append(MovieList[:index], MovieList[index+1:]...)
		}
	}
}

// the "U" of CRUD
func updateMovie(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Endpoint hit (PUT request received): /movie/{id}")

	// Extract the id from the route
	routeVariables := mux.Vars(r)
	updateId := routeVariables["id"]

	// Get the request body
	putBody, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem reading request body")
		return
	}

	// Unmarshal the request body JSON into a new `Movie` struct.
	var movie Movie
	err = json.Unmarshal(putBody, &movie)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem unmarshalling")
	}

	for index, xmovie := range MovieList {
		if xmovie.Id == updateId {
			// Overwrite the existing values
			xmovie.Title = movie.Title
			xmovie.Desc = movie.Desc
			xmovie.ReleaseYear = movie.ReleaseYear

			// Store all movies before the found index, plus the updated movie.
			temporaryList := append(MovieList[:index], xmovie)

			// Update the global MovieList with the updated movie.
			MovieList = append(temporaryList, MovieList[index+1:]...)
		}
	}

	fmt.Fprintf(w, "%+v", string(putBody))
}

func readMoviesFromJSON() []Movie {
	jsonFile, err := os.Open("./data.json")

	if err != nil {
		fmt.Println("error reading json file", err)
	}

	/**
	Defer closing the file, so we can parse it.
	This ƒn will execute after everything else in this ƒn executes - after it
	finishes its' final statement and exits, but before it actually returns.
	Deferred functions are executed in LIFO order.

	Need to call this AFTER checking for an error - because if you get an error,
	you haven't actually gotten the jsonFile. Program would fail
	*/
	defer jsonFile.Close()

	// Read the opened json file as a byte array
	byteValue, _ := io.ReadAll(jsonFile)

	var movies []Movie
	err = json.Unmarshal(byteValue, &movies)

	if err != nil {
		fmt.Println(err.Error(), "\nproblem unmarshalling")
	}

	return movies
}

// var stuffs map[string]interface{}
// The type of the `stuffs` variable = a map where the keys are strings, and the
// values are of type interface{}
