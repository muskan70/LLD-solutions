package models

import (
	"fmt"
	"sync/atomic"
	"time"
)

var movieId atomic.Uint64

type Movie struct {
	Id          uint64
	Title       string
	Description string
	Languages   []int
	Genre       string
	ReleaseDate time.Time
	Duration    time.Duration
	Cast        []uint64
}

func NewMovie(name string) *Movie {
	return &Movie{
		Id:    movieId.Add(1),
		Title: name,
	}
}

func (m *Movie) SetDescription(despt string) {
	m.Description = despt
}

func (m *Movie) SetLanguages(lang []int) {
	m.Languages = append(m.Languages, lang...)
}

func (m *Movie) SetGenre(genre string) {
	m.Genre = genre
}

func (m *Movie) SetReleaseDate(date string) {
	if releaseDate, err := time.Parse("2006-01-02", date); err != nil {
		fmt.Println(err)
	} else {
		m.ReleaseDate = releaseDate
	}
}

func (m *Movie) AddActors(actorIds []uint64) {
	m.Cast = append(m.Cast, actorIds...)
}
