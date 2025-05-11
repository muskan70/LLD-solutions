package services

import (
	"ticketBooking/models"
)

var movieManager *MovieManager

type MovieManager struct {
	Movies map[uint64]*models.Movie
}

func NewMovieManager() {
	movieManager = &MovieManager{
		Movies: make(map[uint64]*models.Movie),
	}
}

func (m *MovieManager) GetMovieById(MovieId uint64) *models.Movie {
	return m.Movies[MovieId]
}

func (m *MovieManager) CreateMovie(name, description string, duration int) {
	movie := models.NewMovie(name)
	m.Movies[movie.Id] = movie
	movie.SetDescription(description)
	movie.SetDuration(duration)
}
