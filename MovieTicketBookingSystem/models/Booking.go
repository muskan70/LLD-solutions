package models

import (
	"sync/atomic"
	"ticketBooking/constants"
)

var bookingId atomic.Uint64

type Booking struct {
	Id            uint64
	ShowId        uint64
	SeatIds       []uint64
	UserId        uint64
	BookingStatus int
}

func NewBooking(showId, userId uint64, seats []uint64) *Booking {
	return &Booking{
		Id:            bookingId.Add(1),
		ShowId:        showId,
		UserId:        userId,
		SeatIds:       seats,
		BookingStatus: constants.BOOKING_STATUS_PENDING,
	}
}

func (b *Booking) ConfirmBooking() {
	b.BookingStatus = constants.BOOKING_STATUS_CONFIRMED
}

func (b *Booking) ExpireBooking() {
	b.BookingStatus = constants.BOOKING_STATUS_FAILED
}
