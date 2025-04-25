package theatre

import "sync/atomic"

var theatreId atomic.Uint64

type Theatre struct {
	Id      uint64
	Name    string
	City    string
	Screens []*Screen
}

func NewTheatre(name string, city string) *Theatre {
	return &Theatre{
		Id:   theatreId.Add(1),
		Name: name,
		City: city,
	}
}

func (t *Theatre) AddScreen(s *Screen) {
	t.Screens = append(t.Screens, s)
}
