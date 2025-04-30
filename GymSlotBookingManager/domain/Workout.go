package domain

import (
	"encoding/json"
	"errors"
	"flip/constants"
	"fmt"
	"sync"
	"sync/atomic"
)

var slotId atomic.Uint64

type Timing struct {
	StartTime int
	EndTime   int
}

type WorkoutSlot struct {
	Id            uint64
	CentreId      uint64
	SlotTime      int
	NumberOfSeats int
	WorkoutType   int
	SlotType      int
	SlotDays      []int
	SeatsBooked   map[uint64]bool
	SeatLock      sync.Mutex
	NormalQueue   []*User
	PremiumQueue  []*User
}

func NewWorkoutSlot(centreId uint64, workoutType, startTime, endTime, availableSeats, slotType int) (*WorkoutSlot, error) {
	if startTime < 0 || startTime > 23 || endTime < 0 || endTime > 23 || endTime <= startTime || endTime-startTime > 1 {
		return nil, errors.New("invalid slot")
	}
	slot := &WorkoutSlot{
		Id:            slotId.Add(1),
		CentreId:      centreId,
		SlotTime:      startTime,
		NumberOfSeats: availableSeats,
		WorkoutType:   workoutType,
		SlotType:      slotType,
		SeatsBooked:   make(map[uint64]bool),
		SeatLock:      sync.Mutex{},
	}
	return slot, nil
}

func (ws *WorkoutSlot) BookSlot(user *User) error {
	availableSeats := ws.NumberOfSeats - len(ws.SeatsBooked)
	if ws.SlotType != constants.SLOT_TYPE_NORMAL && ws.SlotType != user.UserType {
		return errors.New("this user is not eligible for this slot")
	}
	if availableSeats > 0 {
		ws.SeatLock.Lock()
		defer ws.SeatLock.Unlock()
		ws.SeatsBooked[user.UserId] = true
		return nil
	} else if user.UserType == constants.FK_VIP_USER {
		ws.PremiumQueue = append(ws.PremiumQueue, user)
	} else {
		ws.NormalQueue = append(ws.NormalQueue, user)
	}
	return errors.New("slot fully booked, user added to waiting list")
}

func (ws *WorkoutSlot) GetNextUserFromQueue() *User {
	var nextUser *User
	if len(ws.PremiumQueue) > 0 {
		nextUser = ws.PremiumQueue[0]
		ws.PremiumQueue = ws.PremiumQueue[1:]
	} else if len(ws.NormalQueue) > 0 {
		nextUser = ws.NormalQueue[0]
		ws.NormalQueue = ws.NormalQueue[1:]
	}
	return nextUser
}

func (ws *WorkoutSlot) CancelSlot(user *User, day int) {
	ws.SeatLock.Lock()
	defer ws.SeatLock.Unlock()
	delete(ws.SeatsBooked, user.UserId)
	user.RemoveBookedSlot(ws, day)

	workoutDetails, _ := json.Marshal(ws)
	fmt.Println("Workout session cancelled by user", user.UserId, string(workoutDetails))

	for nextUser := ws.GetNextUserFromQueue(); nextUser != nil; {
		if nextUser.CheckDaySlots(day) == nil {
			ws.SeatsBooked[nextUser.UserId] = true
			nextUser.AddBookedSlot(ws, day)
			workoutDetails, _ := json.Marshal(ws)
			fmt.Println("Slot booking confirmed for ", nextUser.UserId, string(workoutDetails))
			break
		}
	}
}
