package main

import "errors"

type User struct {
	Name                string
	BookedActivitySlots []WorkoutSession
}

func registerUser(name string) error {
	if len(name) == 0 {
		return errors.New("invalid name")
	}
	user := User{Name: name}
	users[name] = &user
	return nil
}

func bookSession(userName, centreName, workoutName string, startTime, endTime int) error {
	_, ok := users[userName]
	if !ok {
		return errors.New("user is not registered")
	}
	centre, ok := centres[centreName]
	if !ok {
		return errors.New("invalid centre")
	}
	workoutSessions, ok := centre.DaySchedule[workoutName]
	if !ok {
		return errors.New("invalid workoutName")
	}
	for i, slot := range workoutSessions.Slots {
		if slot.StartTime == startTime && slot.EndTime == endTime {
			if slot.NumberOfSeats >= 1 {
				centres[centreName].DaySchedule[workoutName].Slots[i].NumberOfSeats--
				users[userName].BookedActivitySlots = append(users[userName].BookedActivitySlots, WorkoutSession{
					CentreName:  centreName,
					WorkoutName: workoutName,
					Slots:       []Slot{{StartTime: startTime, EndTime: endTime}},
				})
				return nil
			} else {
				return errors.New("all seats are booked")
			}
		}
	}
	return errors.New("invalid slot")

}
