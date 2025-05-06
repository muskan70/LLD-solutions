package constants

import "time"

const (
	FK_VIP_USER = iota + 1
	FK_NORMAL_USER
)

const (
	SLOT_TYPE_PREMIUM = 1
	SLOT_TYPE_NORMAL  = 2
)

const (
	WorkoutType_CARDIO = iota + 1
	WorkoutType_YOGA
	WorkoutType_SWIMMING
	WorkoutType_WEIGHTS
)

const (
	Booking_Confirmed = iota + 1
	Booking_Cancelled
	Booking_Waiting
	Booking_New
)

var workoutTypes []int

func GetWorkoutTypes() []int {
	workoutTypes = []int{1, 2, 3, 4}
	return workoutTypes
}

const (
	MONDAY = iota + 1
	TUESDAY
	WEDNESDAY
	THRUSDAY
	FRIDAY
	SATURDAY
	SUNDAY
)

func GetWeekday(date string) int {
	curDate, _ := time.Parse("2006-01-02 15:04:05", date+"  00:00:00")
	return int(curDate.Weekday())
}

func GetTodayDate() string {
	return time.Now().Format("2006-01-02")
}
