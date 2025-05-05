package model

import (
	"errors"
	"sync/atomic"
)

var slotId atomic.Uint64

type WorkoutSlot struct {
	Id            uint64
	CentreId      uint64
	SlotTime      int
	NumberOfSeats int
	WorkoutType   int
	SlotType      int
	SlotDays      []int
	Charges       float64
}

func NewWorkoutSlot(centreId uint64, workoutType, slotTime, availableSeats, slotType int, days []int, charges float64) (uint64, error) {
	if slotTime < 0 || slotTime > 23 {
		return 0, errors.New("invalid slot time")
	}
	if len(days) == 0 {
		return 0, errors.New("invalid slot days")
	}
	slot := &WorkoutSlot{
		Id:            slotId.Add(1),
		CentreId:      centreId,
		SlotTime:      slotTime,
		NumberOfSeats: availableSeats,
		WorkoutType:   workoutType,
		SlotType:      slotType,
		SlotDays:      days,
		Charges:       charges,
	}
	WorkoutSlots[slot.Id] = slot
	return slot.Id, nil
}

func GetWorkSlotById(id uint64) (*WorkoutSlot, error) {
	_, ok := WorkoutSlots[id]
	if !ok {
		return nil, errors.New("invalid workout slot")
	}
	return WorkoutSlots[id], nil
}
