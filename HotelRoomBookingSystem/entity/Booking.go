package entity

import (
	"sliceMC/constants"
	"sync/atomic"
	"time"
)

var bookingId atomic.Uint64

type Booking struct {
	Id        uint64
	RoomId    uint64
	UserId    uint64
	Status    int
	Timestamp time.Time
}

func NewBooking(userId, roomId uint64) *Booking {
	return &Booking{
		Id:        bookingId.Add(1),
		RoomId:    roomId,
		UserId:    userId,
		Status:    constants.BOOKING_STATUS_NEW,
		Timestamp: time.Now(),
	}
}
