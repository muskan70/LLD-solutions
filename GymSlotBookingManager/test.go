package main

import (
	"flip/constants"
	"flip/model"
	"flip/usecase"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

func GetAvailableSessionsAndBookRandomSession(userId uint64, workoutType int, date string, wg *sync.WaitGroup) {
	workouts, err := usecase.GetAvailableSessions(userId, workoutType, date)
	if err != nil {
		fmt.Println("Available sessions Error for userId:", userId, "->", err.Error())
		wg.Done()
		return
	} else {
		sessionIds := []uint64{}
		for i := range workouts {
			sessionIds = append(sessionIds, workouts[i].Session.Id)
		}
		fmt.Println("Workout Sessions for ", userId, ":", sessionIds)
	}
	idx := rand.Intn(len(workouts))
	if bookingId, err := usecase.BookSession(userId, workouts[idx].Session); err != nil {
		fmt.Println("Booking Session Error for userId:", userId, "->", err.Error())
	} else {
		fmt.Println("sessionId:", workouts[idx].Session.Id, " booked successfully by userId:", userId, "with bookingId:", bookingId)
	}
	wg.Done()
}

func test() {
	// Create centre1
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

	//Add slots in centreId1 -> 1 : normal slot, 2 : premium slot
	slotId1, err := usecase.CreateWorkoutSlot(centreId1, constants.WorkoutType_WEIGHTS, 6, 1, constants.SLOT_TYPE_NORMAL, []int{constants.MONDAY, constants.TUESDAY, constants.WEDNESDAY, constants.THRUSDAY, constants.FRIDAY}, 100)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId1)
	}

	slotId2, err := usecase.CreateWorkoutSlot(centreId1, constants.WorkoutType_WEIGHTS, 7, 1, constants.SLOT_TYPE_PREMIUM, []int{constants.MONDAY, constants.TUESDAY, constants.THRUSDAY}, 150)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId2)
	}

	// Create centre2
	centreId2, err := model.NewCentre("Kasturinagar", model.NewLocation(4, 4))
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("CentreId Created:", centreId2)
	}

	centre, _ = model.GetCentre(centreId2)
	timings = []model.Timing{
		{StartTime: 6, EndTime: 9},
		{StartTime: 17, EndTime: 21},
	}
	if err := centre.AddTimings(timings); err != nil {
		fmt.Println(err.Error())
	}

	if err := centre.AddWorkoutTypes([]int{constants.WorkoutType_WEIGHTS, constants.WorkoutType_CARDIO, constants.WorkoutType_YOGA}); err != nil {
		fmt.Println(err.Error())
	}

	//Add slots in centreId1 -> 1 : normal slot, 2 : premium slot, 3: premium slot
	slotId1, err = usecase.CreateWorkoutSlot(centreId2, constants.WorkoutType_WEIGHTS, 6, 1, constants.SLOT_TYPE_NORMAL, []int{constants.MONDAY, constants.TUESDAY, constants.WEDNESDAY, constants.FRIDAY}, 100)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId1)
	}

	slotId2, err = usecase.CreateWorkoutSlot(centreId2, constants.WorkoutType_WEIGHTS, 7, 1, constants.SLOT_TYPE_PREMIUM, []int{constants.MONDAY, constants.TUESDAY, constants.THRUSDAY}, 150)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId2)
	}

	slotId3, err := usecase.CreateWorkoutSlot(centreId2, constants.WorkoutType_YOGA, 7, 1, constants.SLOT_TYPE_PREMIUM, []int{constants.MONDAY, constants.TUESDAY, constants.FRIDAY}, 150)
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("SlotId Created:", slotId3)
	}

	model.FillSessionsForWeek(centreId1)
	model.FillSessionsForWeek(centreId2)

	for i := 0; i < 3; i++ {
		userId, err := model.NewUser(strconv.Itoa(97+i)+"xyz", 9678932421, model.NewLocation(float64(i+1), float64(i+2)), constants.FK_VIP_USER)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("UserId created:", userId)
		}
	}

	for i := 0; i < 3; i++ {
		userId, err := model.NewUser(strconv.Itoa(97+i+5)+"xyz", 9678932421, model.NewLocation(float64(i+2), float64(i+1)), constants.FK_NORMAL_USER)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("UserId created:", userId)
		}
	}

	wg := new(sync.WaitGroup)

	for i := 1; i <= 6; i++ {
		wg.Add(1)
		go GetAvailableSessionsAndBookRandomSession(uint64(i), constants.WorkoutType_WEIGHTS, constants.GetTodayDate(), wg)
	}

	wg.Wait()

	for i := 1; i <= 6; i++ {
		if err := usecase.CancelNextSession(uint64(i)); err != nil {
			fmt.Println("userId:", i, err)
		}
	}

}
