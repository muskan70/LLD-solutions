package constants

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
