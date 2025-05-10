package models

import (
	"sync/atomic"
	"ticketBooking/constants"
)

var screenId atomic.Uint64

type Screen struct {
	Id          uint64
	TheatreId   uint64
	Name        string
	NoofRows    int
	NoofColumns int
	Seats       [][]*Seat
}

func NewScreen(name string, theatreId uint64, rows, columns int) *Screen {
	screen := &Screen{
		Id:          screenId.Add(1),
		TheatreId:   theatreId,
		Name:        name,
		NoofRows:    rows,
		NoofColumns: columns,
		Seats:       make([][]*Seat, rows),
	}
	screen.SetSeatsLayout(rows, columns)
	return screen
}

func (s *Screen) SetSeatsLayout(rows, columns int) {
	for i := range rows {
		s.Seats[i] = make([]*Seat, columns)
		for j := range columns {
			if i == rows-1 {
				s.Seats[i][j] = NewSeat(i+1, j+1, constants.SEAT_CATEGORY_PLATINUM)
			} else {
				s.Seats[i][j] = NewSeat(i+1, j+1, constants.SEAT_CATEGORY_SILVER)
			}
		}
	}
}

func (s *Screen) GetSeatsLayout() [][]*Seat {
	return s.Seats
}
