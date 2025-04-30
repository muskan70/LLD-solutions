package usecase

import (
	"errors"
	"flip/constants"
	"flip/domain"
)

var Centres map[uint64]*domain.Centre

type AvailableSlotsWrtUser struct {
	Distance float64
	Slot     *domain.WorkoutSlot
	CentreId uint64
}

func NewCentreUsecase() {
	Centres = make(map[uint64]*domain.Centre)
}

func AddCentre(name string, loc domain.Location) (uint64, error) {
	centre, err := domain.NewCentre(name, loc)
	if err != nil {
		return 0, err
	}
	Centres[centre.ID] = centre
	return centre.ID, nil
}

func AddCentreTimings(centreId uint64, timings []domain.Timing) error {
	if centre, ok := Centres[centreId]; ok {
		return centre.AddTimings(timings)
	}

	return errors.New("Invalid centreId")
}

func AddCentreWorkoutTypes(centreId uint64, activities []int) error {
	if centre, ok := Centres[centreId]; ok {
		return centre.AddWorkoutTypes(activities)
	}
	return errors.New("Invalid centreId")
}

func AddCentreWorkoutSlot(centreId uint64, w *domain.WorkoutSlot, day []int) error {
	if centre, ok := Centres[centreId]; ok {
		return centre.AddWorkoutSlot(day, w)
	}
	return errors.New("Invalid centreId")
}

func ViewAvailableWorkoutSessions(workoutType int, day int, user *domain.User) []*AvailableSlotsWrtUser {
	var availableSessions []*AvailableSlotsWrtUser
	for _, centre := range Centres {
		workoutSchedule, ok := centre.WorkoutTypesDayWiseSchedule[workoutType]
		if !ok {
			break
		}
		if slots, ok := workoutSchedule[day]; ok {
			for x := range slots {
				if slots[x].SlotType == constants.SLOT_TYPE_NORMAL || slots[x].SlotType == user.UserType {
					availableSessions = append(availableSessions, &AvailableSlotsWrtUser{
						Distance: domain.GetDistance(centre.Location, user.Location),
						Slot:     slots[x],
						CentreId: centre.ID,
					})
				}
			}
		}
	}
	return availableSessions
}
