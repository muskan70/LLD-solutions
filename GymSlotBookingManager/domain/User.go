package domain

import (
	"errors"
	"sync/atomic"
)

var userId atomic.Uint64

type User struct {
	UserId             uint64
	Name               string
	Phone              int
	Location           Location
	UserType           int
	DaywiseBookedSlots map[int][]*WorkoutSlot
}

func NewUser(name string, phone int, loc Location, userType int) (*User, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid name")
	}
	if phone == 0 {
		return nil, errors.New("invalid phone")
	}
	if loc.XCoord == 0 || loc.YCoord == 0 {
		return nil, errors.New("Invalid user coordinates")
	}
	user := &User{
		UserId:             userId.Add(1),
		Name:               name,
		Phone:              phone,
		Location:           loc,
		UserType:           userType,
		DaywiseBookedSlots: make(map[int][]*WorkoutSlot),
	}
	return user, nil
}

func (u *User) CheckDaySlots(day int) error {
	if len(u.DaywiseBookedSlots[day]) == 3 {
		return errors.New("user has already booked 3 slots for the day")
	}
	return nil
}

func (u *User) AddBookedSlot(ws *WorkoutSlot, day int) error {
	u.DaywiseBookedSlots[day] = append(u.DaywiseBookedSlots[day], ws)
	return nil
}

func (u *User) RemoveBookedSlot(ws *WorkoutSlot, day int) {
	for i, slot := range u.DaywiseBookedSlots[day] {
		if slot.SlotTime == ws.SlotTime {
			u.DaywiseBookedSlots[day] = append(u.DaywiseBookedSlots[day][:i], u.DaywiseBookedSlots[day][i+1:]...)
		}
	}
}

func (u *User) GetAllBookedSlots(day int) []*WorkoutSlot {
	return u.DaywiseBookedSlots[day]
}
