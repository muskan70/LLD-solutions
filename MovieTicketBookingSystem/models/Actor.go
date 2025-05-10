package models

import "sync/atomic"

var actorId atomic.Uint64

type Actor struct {
	ID          uint64
	Name        string
	Description string
}

func NewActor(name, description string) *Actor {
	return &Actor{
		ID:          actorId.Add(1),
		Name:        name,
		Description: description,
	}
}
