package main

import (
	"net/http"
	"encoding/json"
	"database/sql"
	"log"

	"movies/database"

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
	w.Write(moviesJson)

}