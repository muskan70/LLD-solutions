package usecase

import (
	"flip/domain"
)

var Slots map[uint64]*domain.WorkoutSlot

func NewSlotUsecase() {
	Slots = make(map[uint64]*domain.WorkoutSlot)
}

func CreateWorkoutSlot(centreId uint64, workoutType int, startTime, endTime, availableSeats, slotType int, day []int) (uint64, error) {
	slot, err := domain.NewWorkoutSlot(centreId, workoutType, startTime, endTime, availableSeats, slotType)
	if err != nil {
		return 0, err
	}

	if err := AddCentreWorkoutSlot(centreId, slot, day); err != nil {
		return 0, err
	}

	Slots[slot.Id] = slot
	return slot.Id, nil
}
