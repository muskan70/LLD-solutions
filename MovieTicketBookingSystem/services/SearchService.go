package services

import (
	"slices"
	"time"
)

var searchService *SearchCatalogue

type SearchCatalogue struct {
	MovieManager   *MovieManager
	TheatreManager *TheatreManager
	ShowManager    *ShowManager
}

func NewSearchService(m *MovieManager, t *TheatreManager, s *ShowManager) *SearchCatalogue {
	searchService := &SearchCatalogue{
		MovieManager:   m,
		TheatreManager: t,
		ShowManager:    s,
	}
	return searchService
}

type SearchParams struct {
	City     string     // mandatory
	Genre    *int       // optional
	Language *int       // optional
	Title    *string    // optional
	Date     *time.Time // optional
}

func (s *SearchParams) IsSatisfiedBy(movieId uint64) bool {
	movie := searchService.MovieManager.GetMovieById(movieId)
	if s.Genre != nil && movie.Genre != (*s.Genre) {
		return false
	}
	if s.Language != nil && !slices.Contains(movie.Languages, *s.Language) {
		return false
	}
	if s.Title != nil && movie.Title != (*s.Title) {
		return false
	}
	if s.Date != nil && !movie.ReleaseDate.Before(*s.Date) {
		return false
	}
	return true
}

func (s *SearchCatalogue) SearchMovie(scrh *SearchParams) map[uint64][]uint64 {
	movieShows := make(map[uint64][]uint64)
	for _, showId := range s.ShowManager.showsByCity[scrh.City] {
		movieId := s.ShowManager.GetMovieIdByShowId(showId)
		if scrh.IsSatisfiedBy(movieId) {
			movieShows[movieId] = append(movieShows[movieId], showId)
		}
	}
	return movieShows
}
