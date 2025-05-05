package model

var Bookings map[uint64]*Booking
var Centres map[uint64]*Centre
var Users map[uint64]*User
var WorkoutSlots map[uint64]*WorkoutSlot
var WorkoutSchedule map[string]map[int][]uint64
var WorkoutSessions map[uint64]*WorkoutSession

func Init() {
	Users = make(map[uint64]*User)
	Centres = make(map[uint64]*Centre)
	WorkoutSlots = make(map[uint64]*WorkoutSlot)
	Bookings = make(map[uint64]*Booking)
	WorkoutSchedule = make(map[string]map[int][]uint64)
	WorkoutSessions = make(map[uint64]*WorkoutSession)

}
