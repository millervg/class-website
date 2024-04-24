package main

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"

	"movies/database"
	"movies/models"

	"github.com/gorilla/mux"
	 _ "github.com/mattn/go-sqlite3"
)

var moviesDatabase *sql.DB

func main() {
	router := mux.NewRouter()

	var err error
	moviesDatabase, err = sql.Open("sqlite3", "../database/movies.db")
	if err != nil {
		log.Fatal(err)
	}
	defer moviesDatabase.Close()

	router.HandleFunc("/", hello).Methods(http.MethodPost)
	router.HandleFunc("/movies", GetAllMoviesHandler).Methods(http.MethodPost)
	router.HandleFunc("/addmovie", InsertNewMovieHandler).Methods(http.MethodPost)


	http.ListenAndServe(":8000", router)
}

func hello(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hey!"))
}

func GetAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	movies := moviesdb.GetAllMovies(moviesDatabase)

	moviesJson, err := json.Marshal(movies)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	w.Write(moviesJson)

}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}

func InsertNewMovieHandler(w http.ResponseWriter, r *http.Request) {
	var movie models.Movies

	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	newMovie := moviesdb.InsertNewMovie(moviesDatabase, movie)

	newMovieJson, err := json.Marshal(newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	enableCors(&w)
	w.Write(newMovieJson)
}

