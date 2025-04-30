package main

import (
	"encoding/json"
	"flip/constants"
	"flip/domain"
	"flip/usecase"
	"fmt"
	"math/rand"
	"sync"
)

func GetAvailableSessionsAndBookRandomSession(userId uint64, workoutType, day int, wg *sync.WaitGroup) {
	workouts, err := usecase.GetAvailableSessions(userId, workoutType, day)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Print("workout Slots for ", userId, ":")
		for i := range workouts {
			fmt.Print(" ", workouts[i].Slot.Id)
		}
		fmt.Println()
	}
	idx := rand.Intn(len(workouts))
	if err := usecase.BookSession(userId, workouts[idx].CentreId, workouts[idx].Slot, constants.MONDAY); err != nil {
		fmt.Println(userId, err.Error())
	} else {
		workoutDetails, _ := json.Marshal(workouts[idx])
		fmt.Println("session booked successfully by", userId, string(workoutDetails))
	}
	fmt.Println()
	wg.Done()
}

func test() {
	centreId1, err := usecase.AddCentre("Kormangalam", domain.NewLocation(2, 2))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("CentreId Created:", centreId1)
	}
	timings := []domain.Timing{
		{StartTime: 6, EndTime: 9},
		{StartTime: 18, EndTime: 20},
	}
	if err := usecase.AddCentreTimings(centreId1, timings); err != nil {
		fmt.Println(err.Error())
	}

	if err := usecase.AddCentreWorkoutTypes(centreId1, []int{constants.WorkoutType_WEIGHTS, constants.WorkoutType_CARDIO}); err != nil {
		fmt.Println(err.Error())
	}

	slotId1, err := usecase.CreateWorkoutSlot(centreId1, constants.WorkoutType_WEIGHTS, 6, 7, 1, constants.SLOT_TYPE_NORMAL, []int{constants.MONDAY, constants.WEDNESDAY, constants.FRIDAY})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId1)
	}

	slotId2, err := usecase.CreateWorkoutSlot(centreId1, constants.WorkoutType_WEIGHTS, 7, 8, 1, constants.SLOT_TYPE_PREMIUM, []int{constants.MONDAY, constants.THRUSDAY})
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId2)
	}

	userId1, err := usecase.RegisterUser("muskan", 9678932421, domain.NewLocation(1, 1), constants.FK_VIP_USER)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserId created:", userId1)
	}

	userId2, err := usecase.RegisterUser("vipul", 9756397683, domain.NewLocation(4, 6), constants.FK_NORMAL_USER)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserId created:", userId2)
	}

	userId3, err := usecase.RegisterUser("yash", 9999888877, domain.NewLocation(2, 3), constants.FK_VIP_USER)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserId created:", userId3)
	}
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go GetAvailableSessionsAndBookRandomSession(userId1, constants.WorkoutType_WEIGHTS, constants.MONDAY, wg)
	go GetAvailableSessionsAndBookRandomSession(userId2, constants.WorkoutType_WEIGHTS, constants.MONDAY, wg)
	go GetAvailableSessionsAndBookRandomSession(userId3, constants.WorkoutType_WEIGHTS, constants.MONDAY, wg)

	wg.Wait()

	if err := usecase.CancelSession(userId1, constants.MONDAY); err != nil {
		fmt.Println(userId1, err)
	}
	if err := usecase.CancelSession(userId2, constants.MONDAY); err != nil {
		fmt.Println(userId2, err)
	}
	if err := usecase.CancelSession(userId3, constants.MONDAY); err != nil {
		fmt.Println(userId3, err)
	}

}
