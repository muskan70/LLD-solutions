package main

import (
	"errors"
	"fmt"
)

var centres map[string]*Centre
var users map[string]*User

type Timing struct {
	StartTime int
	EndTime   int
}

type Slot struct {
	StartTime     int
	EndTime       int
	NumberOfSeats int
}

type WorkoutSession struct {
	CentreName  string
	WorkoutName string
	Slots       []Slot
}

func addCentreActivities(name string, activities []string) error {
	if len(activities) == 0 {
		return errors.New("no activities present")
	}
	for _, a := range activities {
		if len(a) == 0 {
			return errors.New("invalid activities")
		}
	}
	centres[name].WorkoutActivities = activities
	return nil
}

func main() {
	centres = make(map[string]*Centre)
	users = make(map[string]*User)
	if err := addCentre("Kormangalam"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(centres)

	timings := []Timing{
		{StartTime: 6, EndTime: 9},
		{StartTime: 18, EndTime: 20},
	}
	if err := addCentreTimings("Kormangalam", timings); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(centres["Kormangalam"])

	if err := addCentreActivities("Kormangalam", []string{"weights", "cardio"}); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(centres["Kormangalam"])
	if err := addWorkoutSessionToCentre("Kormangalam", "weights", 6, 7, 1); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(centres["Kormangalam"])
	if err := registerUser("muskan"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)
	workouts := viewWorkoutAvailability("weights")
	fmt.Println(workouts)
	if err := bookSession("muskan", "Kormangalam", "weights", 6, 7); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users["muskan"])
	fmt.Println(centres["Kormangalam"])

	if err := registerUser("muskan"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)
	//workouts = viewWorkoutAvailability("weights")
	//fmt.Println(workouts)
	if err := bookSession("muskan", "Kormangalam", "weights", 6, 7); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users["muskan"])
	fmt.Println(centres["Kormangalam"])
	/*if err := registerUser("vipul"); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users)
	workouts = viewWorkoutAvailability("weights")
	fmt.Println(workouts)
	if err := bookSession("vipul", "Kormangalam", "weights", 6, 7); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users["vipul"])
	fmt.Println(centres["Kormangalam"])*/
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

func createWorkoutSlot(startTime, endTime, availableSeats int) Slot {
	return Slot{
		StartTime:     startTime,
		EndTime:       endTime,
		NumberOfSeats: availableSeats,
	}
}

func addWorkoutSessionToCentre(centreName, workoutName string, startTime, endTime, availableSeats int) error {
	centre := centres[centreName]
	flag := false
	for _, activity := range centre.WorkoutActivities {
		if workoutName == activity {
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("invalid activity")
	}

	if startTime < 0 {
		return errors.New("invalid slot:a")
	} else if startTime > 24 {
		return errors.New("invalid slot:b")
	} else if endTime < 0 {
		return errors.New("invalid slot:c")
	} else if endTime > 24 {
		return errors.New("invalid slot:d")
	} else if endTime <= startTime {
		return errors.New("invalid slot:e")
	} else if endTime-startTime > 1 {
		return errors.New("invalid slot:f")
	}

	flag = false
	for _, timing := range centre.CentreTimings {
		if startTime >= timing.StartTime && endTime <= timing.EndTime {
			flag = true
			break
		}
	}
	if !flag {
		return errors.New("invalid slot")
	}
	workoutSlot := createWorkoutSlot(startTime, endTime, availableSeats)
	if workout, ok := centres[centreName].DaySchedule[workoutName]; !ok {
		centres[centreName].DaySchedule[workoutName] = WorkoutSession{
			WorkoutName: workoutName,
			CentreName:  centreName,
			Slots:       []Slot{workoutSlot},
		}
	} else {
		workout.Slots = append(workout.Slots, workoutSlot)
		centres[centreName].DaySchedule[workoutName] = workout
	}

	return nil
}

func viewWorkoutAvailability(workoutType string) []WorkoutSession {
	var availableSessions []WorkoutSession
	for _, centre := range centres {
		if workout, ok := centre.DaySchedule[workoutType]; ok {
			availableSessions = append(availableSessions, workout)
		}
	}
	return availableSessions
}
