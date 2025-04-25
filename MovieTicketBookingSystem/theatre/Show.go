package theatre

import (
	"sync/atomic"
	"time"
)

var showId atomic.Uint64

type Show struct {
	Id        uint64
	ScreenId  uint64
	StartTime time.Time
	Duration  int
	MovieId   uint64
}

func NewShow(sceenId, movieId uint64, showTime time.Time) *Show {
	return &Show{
		Id:        showId.Add(1),
		ScreenId:  sceenId,
		MovieId:   movieId,
		StartTime: showTime,
	}
}
