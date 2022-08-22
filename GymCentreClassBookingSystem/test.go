package main

import "fmt"

func test() {
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
	if err := addWorkoutSessionToCentre("Kormangalam", "weights", 6, 7, 100); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(centres["Kormangalam"])
	if err := registerUser("muskan"); err != nil {
		fmt.Println(err.Error())
	}
	workouts := viewWorkoutAvailability("weights")
	fmt.Println(workouts)
	if err := bookSession("muskan", "Kormangalam", "weights", 6, 7); err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(users["muskan"])

}
