package models

import (
	"sync/atomic"
	"ticketBooking/constants"
)

var bookingId atomic.Uint64

type Booking struct {
	Id            uint64
	ShowId        uint64
	SeatIds       []string
	UserId        uint64
	BookingStatus int
}

func NewBooking(showId, userId uint64, seats []string) *Booking {
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

func (b *Booking) CancelBooking() {
	b.BookingStatus = constants.BOOKING_STATUS_CANCELLED
}
