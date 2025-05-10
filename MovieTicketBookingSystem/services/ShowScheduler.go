package services

import (
	"ticketBooking/models"
)

var showManager *ShowManager

type ShowManager struct {
	shows map[uint64]*models.Show
}

func NewShowManager() {
	showManager = &ShowManager{
		shows: make(map[uint64]*models.Show),
	}
}

func (sm *ShowManager) GetShowById(showId uint64) *models.Show {
	return sm.shows[showId]
}

func (sm *ShowManager) CreateShow(movieId uint64, screen *models.Screen, showTime string, prices map[int]float64) {
	show := models.NewShow(screen, movieId, showTime, prices)
	sm.shows[show.Id] = show
}
