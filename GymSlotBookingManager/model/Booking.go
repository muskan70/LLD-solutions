package model

import (
	"errors"
	"flip/constants"
	"sync/atomic"
)

var bookingId atomic.Uint64

type Booking struct {
	Id        uint64
	UserId    uint64
	CentreId  uint64
	SessionId uint64
	Date      string
	Status    int
}

func NewBooking(userId, centreId, sessionId uint64, date string) (*Booking, error) {
	if userId == 0 || sessionId == 0 {
		return nil, errors.New("invalid booking")
	}
	booking := &Booking{
		Id:        bookingId.Add(1),
		UserId:    userId,
		CentreId:  centreId,
		SessionId: sessionId,
		Date:      date,
		Status:    constants.Booking_New,
	}
	Bookings[booking.Id] = booking
	user, _ := GetUserById(userId)
	user.AddBooking(booking.Id, date)
	return booking, nil
}

func GetBookingById(id uint64) (*Booking, error) {
	_, ok := Bookings[id]
	if !ok {
		return nil, errors.New("invalid bookingId")
	}
	return Bookings[id], nil
}

func (b *Booking) NotifyUser() {
	user, _ := GetUserById(b.UserId)
	user.Notify(b)
}

func (b *Booking) UpdateBookingStatus(status int) {
	b.Status = status
	b.NotifyUser()
}
