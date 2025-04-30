package usecase

import (
	"errors"
	"flip/domain"
	"sort"
)

var Users map[uint64]*domain.User

func NewUserUsecase() {
	Users = make(map[uint64]*domain.User)
}

func RegisterUser(name string, phone int, loc domain.Location, usertype int) (uint64, error) {
	user, err := domain.NewUser(name, phone, loc, usertype)
	if err != nil {
		return 0, err
	}
	Users[user.UserId] = user
	return user.UserId, nil
}

func GetAvailableSessions(userId uint64, workoutType int, day int) ([]*AvailableSlotsWrtUser, error) {
	_, ok := Users[userId]
	if !ok {
		return nil, errors.New("user is not registered")
	}
	slots := ViewAvailableWorkoutSessions(workoutType, day, Users[userId])
	sort.Slice(slots, func(i, j int) bool {
		return slots[i].Distance < slots[j].Distance
	})
	return slots, nil
}

func BookSession(userId, centreId uint64, workoutSlot *domain.WorkoutSlot, day int) error {
	user, ok := Users[userId]
	if !ok {
		return errors.New("user is not registered")
	}
	if err := user.CheckDaySlots(day); err != nil {
		return err
	}
	centre, ok := Centres[centreId]
	if !ok {
		return errors.New("invalid centre")
	}

	if schedule, ok := centre.WorkoutTypesDayWiseSchedule[workoutSlot.WorkoutType]; ok {
		if slots, ok := schedule[day]; ok {
			for _, slot := range slots {
				if slot.SlotTime == workoutSlot.SlotTime {
					if err := slot.BookSlot(user); err != nil {
						return err
					}
					user.AddBookedSlot(workoutSlot, day)
					return nil
				}
			}
		}
	}
	return errors.New("invalid slot")
}

func CancelSession(userId uint64, day int) error {
	user, ok := Users[userId]
	if !ok {
		return errors.New("user is not registered")
	}
	ws := user.GetAllBookedSlots(day)
	if ws == nil {
		return errors.New("no booked slots present")
	}
	ws[0].CancelSlot(user, day)
	return nil
}

func GetUserBookedSlotsForDay(userId uint64, day int) ([]*domain.WorkoutSlot, error) {
	user, ok := Users[userId]
	if !ok {
		return nil, errors.New("user is not registered")
	}
	return user.GetAllBookedSlots(day), nil
}
