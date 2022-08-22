package main

import "errors"

func addCentre(name string) error {
	if len(name) == 0 {
		return errors.New("invalid name")
	}
	centre := Centre{Name: name, DaySchedule: make(map[string]WorkoutSession)}
	centres[name] = &centre
	return nil
}

func addCentreTimings(name string, timings []Timing) error {
	if len(timings) == 0 {
		return errors.New("no timings present")
	}
	for _, t := range timings {
		if t.StartTime >= t.EndTime || t.StartTime < 0 || t.StartTime > 24 || t.EndTime < 0 || t.EndTime > 24 {
			return errors.New("invalid timings")
		}
	}
	centres[name].CentreTimings = timings
	return nil
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
