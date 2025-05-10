package models

import (
	"sync"
	"sync/atomic"
	"ticketBooking/constants"
)

var showId atomic.Uint64

type SeatLock struct {
	Status int
	Lock   sync.Mutex
	Price  float64
}

type Show struct {
	Id        uint64
	ScreenId  uint64
	StartTime string
	Duration  int
	MovieId   uint64
	ShowSeats map[string]*SeatLock
}

func NewShow(screen *Screen, movieId uint64, showTime string, prices map[int]float64) *Show {
	show := &Show{
		Id:        showId.Add(1),
		ScreenId:  screen.Id,
		MovieId:   movieId,
		StartTime: showTime,
	}
	show.fillShowSeats(screen, prices)
	return show
}

func (s *Show) fillShowSeats(screen *Screen, prices map[int]float64) {
	seatsMap := make(map[string]*SeatLock)
	seats := screen.GetSeatsLayout()
	for i := range seats {
		for j := range seats[i] {
			price := prices[constants.SEAT_CATEGORY_PLATINUM]
			if seats[i][j].SeatCategory == constants.SEAT_CATEGORY_SILVER {
				price = prices[constants.SEAT_CATEGORY_SILVER]
			}
			seatsMap[seats[i][j].Id] = &SeatLock{
				Status: constants.SEAT_STATUS_AVAILABLE,
				Lock:   sync.Mutex{},
				Price:  price,
			}
		}
	}
	s.ShowSeats = seatsMap
}

func (s *Show) CheckSeatAvailability(seatId string) bool {
	sL, ok := s.ShowSeats[seatId]
	return ok && sL.Status == constants.SEAT_STATUS_AVAILABLE
}

func (s *Show) UpdateSeatStatus(seatId string, status int, wg *sync.WaitGroup) {
	if sL, ok := s.ShowSeats[seatId]; ok && sL.Status == constants.SEAT_STATUS_AVAILABLE {
		sL.Lock.Lock()
		defer sL.Lock.Unlock()
		sL.Status = status
	}
}

func (s *Show) GetAvailableSeats() []string {
	var seats []string
	for seatId, seatLock := range s.ShowSeats {
		if seatLock.Status == constants.SEAT_STATUS_AVAILABLE {
			seats = append(seats, seatId)
		}
	}
	return seats
}
