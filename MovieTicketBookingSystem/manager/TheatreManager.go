package manager

import (
	"ticketBooking/models"
)

var theatreManager *TheatreManager

type TheatreManager struct {
	theatres map[uint64]*models.Theatre
	screens  map[uint64]*models.Screen
}

func NewTheatreManager() *TheatreManager {
	theatreManager = &TheatreManager{
		theatres: make(map[uint64]*models.Theatre),
	}
	return theatreManager
}

func (m *TheatreManager) GetTheatreById(id uint64) *models.Theatre {
	return m.theatres[id]
}

func (m *TheatreManager) GetTheatreCityByScreenId(screenId uint64) string {
	screen := m.screens[screenId]
	theatre := m.theatres[screen.TheatreId]
	return theatre.City
}

func (m *TheatreManager) CreateTheatre(name, city, address string) {
	t := models.NewTheatre(name, city, address)
	m.theatres[t.Id] = t
}

func (m *TheatreManager) CreateScreen(name string, theatreId uint64, rows, columns int) {
	scrn := models.NewScreen(name, theatreId, rows, columns)
	m.screens[scrn.Id] = scrn
	theatre := m.theatres[theatreId]
	theatre.AddScreen(scrn.Id)
}
