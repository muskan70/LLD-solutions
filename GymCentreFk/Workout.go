package main

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

func createWorkoutSlot(startTime, endTime, availableSeats int) Slot {
	return Slot{
		StartTime:     startTime,
		EndTime:       endTime,
		NumberOfSeats: availableSeats,
	}
}
