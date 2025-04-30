package domain

import (
	"errors"
	"flip/constants"
	"sort"
	"sync/atomic"
)

var centreId atomic.Uint64

type Centre struct {
	ID                          uint64
	Name                        string
	Location                    Location
	CentreTimings               []Timing
	WorkoutTypesDayWiseSchedule map[int]map[int][]*WorkoutSlot
}

// type CentreBookingSchedulePerDay struct {
// 	Id            uint64
// 	WorkoutSlotId uint64
// 	Date          string
// 	Day           string
// 	SeatsBooked   map[uint64]bool
// 	SeatLock      sync.Mutex
// 	NormalQueue   []*User
// 	PremiumQueue  []*User
// }

func NewCentre(name string, loc Location) (*Centre, error) {
	if len(name) == 0 {
		return nil, errors.New("invalid centre name")
	}

	if loc.XCoord == 0 || loc.YCoord == 0 {
		return nil, errors.New("Invalid centre coordinates")
	}

	centre := &Centre{
		ID:                          centreId.Add(1),
		Name:                        name,
		Location:                    loc,
		WorkoutTypesDayWiseSchedule: make(map[int]map[int][]*WorkoutSlot),
	}
	return centre, nil
}

func (c *Centre) AddTimings(timings []Timing) error {
	if len(timings) == 0 {
		return errors.New("no timings present")
	}
	for _, t := range timings {
		if t.StartTime >= t.EndTime || t.StartTime < 0 || t.StartTime > 24 || t.EndTime < 0 || t.EndTime > 24 {
			return errors.New("invalid timings")
		}
	}
	c.CentreTimings = timings
	return nil
}

func (c *Centre) AddWorkoutTypes(workoutTypes []int) error {
	if len(workoutTypes) == 0 {
		return errors.New("no workout types present")
	}

	constWorkoutTypes := constants.GetWorkoutTypes()
	for _, a := range workoutTypes {
		if sort.SearchInts(constWorkoutTypes, a) == len(constWorkoutTypes) {
			return errors.New("invalid workoutTypes")
		} else {
			c.WorkoutTypesDayWiseSchedule[a] = make(map[int][]*WorkoutSlot)
		}
	}
	return nil
}

func (c *Centre) AddWorkoutSlot(days []int, w *WorkoutSlot) error {
	workoutSchedule, ok := c.WorkoutTypesDayWiseSchedule[w.WorkoutType]
	if !ok {
		return errors.New("this workout type is not present in this centre")
	}

	flag := false
	for _, timing := range c.CentreTimings {
		if w.SlotTime >= timing.StartTime && w.SlotTime+1 <= timing.EndTime {
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("invalid slot timings")
	}
	for _, day := range days {
		schedule, ok := workoutSchedule[day]
		if !ok {
			workoutSchedule[day] = append(workoutSchedule[day], w)
			continue
		}
		for i := range schedule {
			if w.SlotTime == schedule[i].SlotTime {
				return errors.New("this workout slot already exists")
			}
		}
		workoutSchedule[day] = append(workoutSchedule[day], w)
	}
	c.WorkoutTypesDayWiseSchedule[w.WorkoutType] = workoutSchedule
	return nil
}
