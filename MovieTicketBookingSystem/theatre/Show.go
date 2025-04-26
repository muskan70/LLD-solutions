package theatre

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
	SeatsStatusMap map[uint64]*SeatLock
}

func NewShow(screen *Screen, movieId uint64, showTime time.Time) *Show {
	seatsMap := make(map[uint64]*SeatLock)
	seats := screen.GetSeats()
	for i := range seats {
		seatsMap[seats[i].SeatNo] = &SeatLock{
			Status: constants.SEAT_STATUS_AVAILABLE,
			Lock:   sync.Mutex{},
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

func (s *Show) CheckSeatAvailability(seatNo uint64) bool {
	sL, ok := s.SeatsStatusMap[seatNo]
	return ok && sL.Status == constants.SEAT_STATUS_AVAILABLE
}

func (s *Show) UpdateSeatStatus(seatNo uint64, status int) {
	if sL, ok := s.SeatsStatusMap[seatNo]; ok {
		sL.Lock.Lock()
		defer sL.Lock.Unlock()
		sL.Status = status
	}
}

func (s *Show) GetAvailableSeats() []uint64 {
	var seats []uint64
	for seatNo, seatLock := range s.SeatsStatusMap {
		if seatLock.Status == constants.SEAT_STATUS_AVAILABLE {
			seats = append(seats, seatNo)
		}
	}
	return seats
}
