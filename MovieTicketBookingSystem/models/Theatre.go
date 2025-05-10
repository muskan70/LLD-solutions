package models

import "sync/atomic"

var theatreId atomic.Uint64

type Theatre struct {
	Id      uint64
	Name    string
	City    string
	Address string
	Screens []uint64
}

func NewTheatre(name, city, address string) *Theatre {
	return &Theatre{
		Id:      theatreId.Add(1),
		Name:    name,
		City:    city,
		Address: address,
	}
}

func (t *Theatre) AddScreen(s uint64) {
	t.Screens = append(t.Screens, s)
}
