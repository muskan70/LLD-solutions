package main

type Centre struct {
	ID                int
	Name              string
	Location          string
	CentreTimings     []Timing
	WorkoutActivities []string
	DaySchedule       map[string]WorkoutSession
}
