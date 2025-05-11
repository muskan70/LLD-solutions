package services

import (
	"ticketBooking/manager"
)

type SearchCatalogue struct {
	MovieManager   *manager.MovieManager
	TheatreManager *manager.TheatreManager
	ShowManager    *manager.ShowManager
}

func NewSearchService(m *manager.MovieManager, t *manager.TheatreManager, s *manager.ShowManager) *SearchCatalogue {
	return &SearchCatalogue{
		MovieManager:   m,
		TheatreManager: t,
		ShowManager:    s,
	}
}

type SearchParams struct {
	City     string  // mandatory
	Date     string  //manatory
	Genre    *int    // optional
	Language *int    // optional
	Title    *string // optional
}

func (s *SearchParams) IsSatisfiedBy(searchService *SearchCatalogue, movieId uint64, showId uint64) bool {
	movie := searchService.MovieManager.GetMovieById(movieId)
	if s.Title != nil && movie.Title != (*s.Title) {
		return false
	}
	if s.Genre != nil && movie.Genre != (*s.Genre) {
		return false
	}
	show := searchService.ShowManager.GetShowById(showId)
	if s.Language != nil && show.Language != *(s.Language) {
		return false
	}

	return true
}

func (s *SearchCatalogue) SearchMovie(scrh *SearchParams) map[uint64][]uint64 {
	movieShows := make(map[uint64][]uint64)
	for _, showId := range s.ShowManager.GetShowsByCity(scrh.City, scrh.Date) {
		movieId := s.ShowManager.GetMovieIdByShowId(showId)
		if scrh.IsSatisfiedBy(s, movieId, showId) {
			movieShows[movieId] = append(movieShows[movieId], showId)
		}
	}
	return movieShows
}
