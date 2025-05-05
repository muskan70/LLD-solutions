package model

import (
	"errors"
	"flip/constants"
	"sort"
	"sync/atomic"
)

var centreId atomic.Uint64

type Centre struct {
	ID                   uint64
	Name                 string
	Location             Location
	CentreTimings        []Timing
	WorkoutTypesSchedule map[int][]uint64
}

func NewCentre(name string, loc Location) (uint64, error) {
	if len(name) == 0 {
		return 0, errors.New("invalid centre name")
	}

	if !loc.IsValid() {
		return 0, errors.New("Invalid centre coordinates")
	}

	centre := &Centre{
		ID:                   centreId.Add(1),
		Name:                 name,
		Location:             loc,
		WorkoutTypesSchedule: make(map[int][]uint64),
	}
	Centres[centre.ID] = centre
	return centre.ID, nil
}

func GetCentre(id uint64) (*Centre, error) {
	_, ok := Centres[id]
	if !ok {
		return nil, errors.New("invalid centreID")
	}
	return Centres[id], nil
}

func (c *Centre) AddTimings(timings []Timing) error {
	if len(timings) == 0 {
		return errors.New("no timings present")
	}
	for _, t := range timings {
		if !t.IsValid() {
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
			c.WorkoutTypesSchedule[a] = []uint64{}
		}
	}
	return nil
}

func (c *Centre) AddWorkoutSlot(slotId uint64) error {
	ws, err := GetWorkSlotById(slotId)
	if err != nil {
		return err
	}
	schedule, ok := c.WorkoutTypesSchedule[ws.WorkoutType]
	if !ok {
		return errors.New("this workout type is not present in this centre")
	}

	flag := false
	for _, timing := range c.CentreTimings {
		if ws.SlotTime >= timing.StartTime && ws.SlotTime+1 <= timing.EndTime {
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("invalid slot timings")
	}
	for i := range schedule {
		curSlot, _ := GetWorkSlotById(schedule[i])
		if err != nil {
			return err
		}
		if ws.SlotTime == curSlot.SlotTime {
			return errors.New("this workout slot already exists")
		}
	}
	schedule = append(schedule, slotId)
	c.WorkoutTypesSchedule[ws.WorkoutType] = schedule
	return nil
}
