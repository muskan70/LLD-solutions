package main

import "errors"

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

	if startTime < 0 || startTime > 24 || endTime < 0 || endTime > 24 || endTime <= startTime || endTime-startTime > 1 {
		return errors.New("invalid slot")
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
