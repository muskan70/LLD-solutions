package theatre

import "sync/atomic"

var screenId atomic.Uint64

type Screen struct {
	Id    uint64
	Name  string
	Seats []*Seat
}

func NewScreen(name string) *Screen {
	return &Screen{
		Id:   screenId.Add(1),
		Name: name,
	}
}

func (s *Screen) AddSeats(seats []*Seat) {
	s.Seats = append(s.Seats, seats...)
}

func (s *Screen) GetSeats() []*Seat {
	return s.Seats
}
