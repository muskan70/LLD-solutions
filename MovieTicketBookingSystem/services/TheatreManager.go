package services

import (
	"ticketBooking/models"
)

var theatreManager *TheatreManager

type TheatreManager struct {
	Theatres map[uint64]*models.Theatre
	Screens  map[uint64]*models.Screen
}

func NewTheatreManager() *TheatreManager {
	theatreManager = &TheatreManager{
		Theatres: make(map[uint64]*models.Theatre),
	}
	return theatreManager
}

func (m *TheatreManager) GetTheatreById(id uint64) *models.Theatre {
	return m.Theatres[id]
}

func (m *TheatreManager) GetTheatreCityByScreenId(screenId uint64) string {
	screen := m.Screens[screenId]
	theatre := m.Theatres[screen.TheatreId]
	return theatre.City
}

func (m *TheatreManager) CreateTheatre(name, city, address string) {
	t := models.NewTheatre(name, city, address)
	m.Theatres[t.Id] = t
}

func (m *TheatreManager) CreateScreen(name string, theatreId uint64, rows, columns int) {
	scrn := models.NewScreen(name, theatreId, rows, columns)
	m.Screens[scrn.Id] = scrn
	theatre := m.Theatres[theatreId]
	theatre.AddScreen(scrn.Id)
}
