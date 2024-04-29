package models

type Movies struct {
	MovieId int 
	Title string
	Genre string 
	Year int
	ImageName string
}

type WatchedMovies struct {
	WatchedMovieId int
	Title string
	Rating int
	Comment string
}