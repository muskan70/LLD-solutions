package services

import (
	"ticketBooking/models"
)

var showManager *ShowManager

type ShowManager struct {
	TheatreManager *TheatreManager
	shows          map[uint64]*models.Show
	showsByCity    map[string][]uint64
}

func NewShowManager(t *TheatreManager) *ShowManager {
	showManager = &ShowManager{
		TheatreManager: t,
		shows:          make(map[uint64]*models.Show),
		showsByCity:    make(map[string][]uint64),
	}
	return showManager
}

func (sm *ShowManager) GetShowById(showId uint64) *models.Show {
	return sm.shows[showId]
}

func (sm *ShowManager) GetMovieIdByShowId(showId uint64) uint64 {
	return sm.shows[showId].MovieId
}

func (sm *ShowManager) CreateShow(movieId uint64, screen *models.Screen, showTime string, prices map[int]float64) {
	show := models.NewShow(screen, movieId, showTime, prices)
	sm.shows[show.Id] = show
	city := sm.TheatreManager.GetTheatreCityByScreenId(screen.Id)
	sm.showsByCity[city] = append(sm.showsByCity[city], show.Id)
}
