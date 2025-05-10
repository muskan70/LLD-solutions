package models

import (
	"sync"
	"sync/atomic"
	"ticketBooking/constants"
	"time"
)

var showId atomic.Uint64

type SeatLock struct {
	Status int
	Lock   sync.Mutex
}

type Show struct {
	Id             uint64
	ScreenId       uint64
	StartTime      time.Time
	Duration       int
	MovieId        uint64
	SeatsStatusMap map[string]*SeatLock
}

func NewShow(screen *Screen, movieId uint64, showTime time.Time) *Show {
	seatsMap := make(map[string]*SeatLock)
	seats := screen.GetSeatsLayout()
	for i := range seats {
		for j := range seats[i] {
			seatsMap[seats[i][j].Id] = &SeatLock{
				Status: constants.SEAT_STATUS_AVAILABLE,
				Lock:   sync.Mutex{},
			}
		}
	}
	return &Show{
		Id:             showId.Add(1),
		ScreenId:       screen.Id,
		MovieId:        movieId,
		StartTime:      showTime,
		SeatsStatusMap: seatsMap,
	}
}

func (s *Show) CheckSeatAvailability(seatId string) bool {
	sL, ok := s.SeatsStatusMap[seatId]
	return ok && sL.Status == constants.SEAT_STATUS_AVAILABLE
}

func (s *Show) UpdateSeatStatus(seatId string, status int) {
	if sL, ok := s.SeatsStatusMap[seatId]; ok {
		sL.Lock.Lock()
		defer sL.Lock.Unlock()
		sL.Status = status
	}
}

func (s *Show) GetAvailableSeats() []string {
	var seats []string
	for seatId, seatLock := range s.SeatsStatusMap {
		if seatLock.Status == constants.SEAT_STATUS_AVAILABLE {
			seats = append(seats, seatId)
		}
	}
	return seats
}
