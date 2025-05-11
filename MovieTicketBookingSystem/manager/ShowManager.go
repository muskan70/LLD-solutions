package manager

import (
	"ticketBooking/models"
)

var showManager *ShowManager

type ShowManager struct {
	shows            map[uint64]*models.Show
	showsByCityNDate map[string]map[string][]uint64
}

func NewShowManager() *ShowManager {
	showManager = &ShowManager{
		shows:            make(map[uint64]*models.Show),
		showsByCityNDate: make(map[string]map[string][]uint64),
	}
	return showManager
}

func (sm *ShowManager) GetShowById(showId uint64) *models.Show {
	return sm.shows[showId]
}

func (sm *ShowManager) GetMovieIdByShowId(showId uint64) uint64 {
	return sm.shows[showId].MovieId
}

func (sm *ShowManager) GetShowsByCity(city string, date string) []uint64 {
	return sm.showsByCityNDate[city][date]
}

func (sm *ShowManager) CreateShow(movieId, screenId uint64, showTime string, date string, prices map[int]float64, lang int) uint64 {
	screen := theatreManager.GetScreenById(screenId)
	show := models.NewShow(screen, movieId, showTime, date, prices, lang)
	sm.shows[show.Id] = show
	city := theatreManager.GetTheatreCityByScreenId(screenId)
	if _, ok := sm.showsByCityNDate[city]; !ok {
		sm.showsByCityNDate[city] = make(map[string][]uint64)
	}
	sm.showsByCityNDate[city][date] = append(sm.showsByCityNDate[city][date], show.Id)
	return show.Id
}
