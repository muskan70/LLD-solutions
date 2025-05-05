package main

import (
	"flip/constants"
	"flip/model"
	"flip/usecase"
	"fmt"
	"math/rand"
	"sync"
)

func GetAvailableSessionsAndBookRandomSession(userId uint64, workoutType int, date string, wg *sync.WaitGroup) {
	workouts, err := usecase.GetAvailableSessions(userId, workoutType, date)
	if err != nil {
		fmt.Println(err.Error())
		wg.Done()
		return
	} else {
		fmt.Print("workout Slots for ", userId, ":")
		for i := range workouts {
			fmt.Print(" ", workouts[i].Session.Id)
		}
		fmt.Println()
	}
	idx := rand.Intn(len(workouts))
	if bookingId, err := usecase.BookSession(userId, workouts[idx].Session); err != nil {
		fmt.Println(userId, "Error:", err.Error())
	} else {
		fmt.Println("sessionId:", workouts[idx].Session.Id, " booked successfully by userId:", userId, "with bookingId:", bookingId)
	}
	wg.Done()
}

func test() {
	centreId1, err := model.NewCentre("Kormangalam", model.NewLocation(2, 2))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("CentreId Created:", centreId1)
	}

	centre, _ := model.GetCentre(centreId1)
	timings := []model.Timing{
		{StartTime: 6, EndTime: 9},
		{StartTime: 18, EndTime: 20},
	}
	if err := centre.AddTimings(timings); err != nil {
		fmt.Println(err.Error())
	}

	if err := centre.AddWorkoutTypes([]int{constants.WorkoutType_WEIGHTS, constants.WorkoutType_CARDIO}); err != nil {
		fmt.Println(err.Error())
	}

	slotId1, err := usecase.CreateWorkoutSlot(centreId1, constants.WorkoutType_WEIGHTS, 6, 1, constants.SLOT_TYPE_NORMAL, []int{constants.MONDAY, constants.WEDNESDAY, constants.FRIDAY}, 100)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId1)
	}

	slotId2, err := usecase.CreateWorkoutSlot(centreId1, constants.WorkoutType_WEIGHTS, 7, 1, constants.SLOT_TYPE_PREMIUM, []int{constants.MONDAY, constants.THRUSDAY}, 150)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId2)
	}

	model.FillSessionsForWeek(centreId1)

	userId1, err := model.NewUser("muskan", 9678932421, model.NewLocation(1, 1), constants.FK_VIP_USER)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserId created:", userId1)
	}

	userId2, err := model.NewUser("vipul", 9756397683, model.NewLocation(4, 6), constants.FK_NORMAL_USER)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserId created:", userId2)
	}

	userId3, err := model.NewUser("yash", 9999888877, model.NewLocation(2, 3), constants.FK_VIP_USER)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("UserId created:", userId3)
	}
	wg := new(sync.WaitGroup)
	wg.Add(3)

	go GetAvailableSessionsAndBookRandomSession(userId1, constants.WorkoutType_WEIGHTS, "2025-05-05", wg)
	go GetAvailableSessionsAndBookRandomSession(userId2, constants.WorkoutType_WEIGHTS, "2025-05-05", wg)
	go GetAvailableSessionsAndBookRandomSession(userId3, constants.WorkoutType_WEIGHTS, "2025-05-05", wg)

	wg.Wait()

	if err := usecase.CancelNextSession(userId1); err != nil {
		fmt.Println(userId1, err)
	}
	if err := usecase.CancelNextSession(userId2); err != nil {
		fmt.Println(userId2, err)
	}
	if err := usecase.CancelNextSession(userId3); err != nil {
		fmt.Println(userId3, err)
	}

}
