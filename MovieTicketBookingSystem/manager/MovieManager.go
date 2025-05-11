package manager

import (
	"ticketBooking/models"
)

var movieManager *MovieManager

type MovieManager struct {
	movies map[uint64]*models.Movie
}

func NewMovieManager() *MovieManager {
	movieManager = &MovieManager{
		movies: make(map[uint64]*models.Movie),
	}
	return movieManager
}

func (m *MovieManager) GetMovieById(MovieId uint64) *models.Movie {
	return m.movies[MovieId]
}

func (m *MovieManager) CreateMovie(name, description string, duration int, languages []int, genre int, releaseDate string) uint64 {
	movie := models.NewMovie(name)
	m.movies[movie.Id] = movie
	movie.SetDescription(description)
	movie.SetDuration(duration)
	movie.SetLanguages(languages)
	movie.SetReleaseDate(releaseDate)
	movie.SetGenre(genre)
	return movie.Id
}
