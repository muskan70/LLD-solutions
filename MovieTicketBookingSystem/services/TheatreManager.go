package services

import (
	"ticketBooking/models"
)

var theatreManager *TheatreManager

type TheatreManager struct {
	Theatres       map[uint64]*models.Theatre
	Screens        map[uint64]*models.Screen
	TheatresByCity map[string][]uint64
}

func NewManager() {
	theatreManager = &TheatreManager{
		Theatres: make(map[uint64]*models.Theatre),
	}
}

func (m *TheatreManager) GetTheatreById(TheatreId uint64) *models.Theatre {
	return m.Theatres[TheatreId]
}

func (m *TheatreManager) CreateTheatre(name, city, address string) {
	t := models.NewTheatre(name, city, address)
	m.Theatres[t.Id] = t
	m.TheatresByCity[city] = append(m.TheatresByCity[city], t.Id)
}

func (m *TheatreManager) CreateScreen(name string, theatreId uint64, rows, columns int) {
	scrn := models.NewScreen(name, theatreId, rows, columns)
	m.Screens[scrn.Id] = scrn
	theatre := m.Theatres[theatreId]
	theatre.AddScreen(scrn.Id)
}
