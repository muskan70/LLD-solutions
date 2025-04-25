package theatre

import "sync/atomic"

var screenId atomic.Uint64

type Screen struct {
	Id    uint64
	Name  string
	Seats []Seat
}

func NewScreen(name string) *Screen {
	return &Screen{
		Id:   screenId.Add(1),
		Name: name,
	}
}
