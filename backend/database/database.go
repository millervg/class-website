package moviesdb

import (
	"database/sql"
	"log"
	"movies/models"
)

func GetAllMovies(db *sql.DB) []models.Movies { 
	var movies[] models.Movies 

	row, err := db.Query("SELECT * FROM movies")
	if err != nil {
		log.Fatal(err)
	}
	defer row.Close()

	for row.Next() {
		var movieId int
		var title string
		var genre string
		var year int
		var imageName string
		row.Scan(&movieId, &title, &genre, &year, &imageName)
		movie := models.Movies {
			MovieId: movieId,
			Title: title,
			Genre: genre,
			Year: year,
			ImageName: imageName,
		}
		movies = append(movies, movie) 
	}
	return movies
}

func InsertNewMovie(db *sql.DB, movie models.Movies) models.Movies {
	insertMovieSql := "INSERT INTO movies (title, genre, year, imageName) VALUES (?, ?, ?)"
	statement, err := db.Prepare(insertMovieSql)
	if err != nil {
		log.Fatal(err)
	}
	res, err := statement.Exec(movie.Title, movie.Genre, movie.Year, movie.ImageName)
	if err != nil {
		log.Fatal(err)
	}
	lid, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	movie.MovieId = int(lid)
	return movie
}