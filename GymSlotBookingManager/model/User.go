package model

import (
	"errors"
	"flip/constants"
	"fmt"
	"sync/atomic"
)

var userId atomic.Uint64

type User struct {
	UserId              uint64
	Name                string
	Phone               int
	Location            Location
	UserType            int
	DatewiseBookedSlots map[string][]uint64
}

func NewUser(name string, phone int, loc Location, userType int) (uint64, error) {
	if len(name) == 0 {
		return 0, errors.New("invalid name")
	}
	if phone == 0 {
		return 0, errors.New("invalid phone")
	}
	if !loc.IsValid() {
		return 0, errors.New("invalid user coordinates")
	}
	user := &User{
		UserId:              userId.Add(1),
		Name:                name,
		Phone:               phone,
		Location:            loc,
		UserType:            userType,
		DatewiseBookedSlots: make(map[string][]uint64),
	}
	Users[user.UserId] = user
	return user.UserId, nil
}

func GetUserById(id uint64) (*User, error) {
	_, ok := Users[id]
	if !ok {
		return nil, errors.New("user is not registered")
	}
	return Users[id], nil
}

func (u *User) CheckDateSlots(date string, centreId uint64) error {
	count := 0
	for _, bookingId := range u.DatewiseBookedSlots[date] {
		booking, _ := GetBookingById(bookingId)
		if booking.CentreId == centreId && booking.Status == constants.Booking_Confirmed {
			count++
		}
	}
	if count >= 3 {
		return errors.New("user has already booked 3 slots for the given date")
	}
	return nil
}

func (u *User) AddBooking(bookingId uint64, date string) error {
	u.DatewiseBookedSlots[date] = append(u.DatewiseBookedSlots[date], bookingId)
	return nil
}

func (u *User) GetAllConfirmedBookingsByDate(date string) []*Booking {
	bookings := []*Booking{}
	for _, bookingId := range u.DatewiseBookedSlots[date] {
		booking, _ := GetBookingById(bookingId)
		if booking.Status == constants.Booking_Confirmed {
			bookings = append(bookings, booking)
		}
	}
	return bookings
}

func (u *User) GetAllBookings() map[string][]uint64 {
	return u.DatewiseBookedSlots
}

func (u *User) Notify(booking *Booking) {
	if booking.Status == constants.Booking_Confirmed {
		fmt.Println(u.UserId, ":Your booking have been confirmed")
	} else if booking.Status == constants.Booking_Waiting {
		fmt.Println(u.UserId, ":Your booking added to waiting list")
	} else if booking.Status == constants.Booking_Cancelled {
		fmt.Println(u.UserId, ":Your booking have been cancelled")
	}
}
