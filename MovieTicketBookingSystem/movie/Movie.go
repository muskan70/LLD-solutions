package movie

import (
	"fmt"
	"sync/atomic"
	"time"
)

var movieId atomic.Uint64

type Movie struct {
	Id          uint64
	Name        string
	Description string
	Language    string
	Genre       string
	ReleaseDate time.Time
	Duration    time.Duration
	Actors      []Actor
}

func NewMovie(name string) *Movie {
	return &Movie{
		Id:   movieId.Add(1),
		Name: name,
	}
}

func (m *Movie) SetDescription(despt string) {
	m.Description = despt
}

func (m *Movie) SetLanguage(lang string) {
	m.Language = lang
}

func (m *Movie) SetGenre(genre string) {
	m.Language = genre
}

func (m *Movie) SetReleaseDate(date string) {
	if releaseDate, err := time.Parse("2006-01-02", date); err != nil {
		fmt.Println(err)
	} else {
		m.ReleaseDate = releaseDate
	}
}
