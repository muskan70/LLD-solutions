package model

import (
	"encoding/json"
	"errors"
	"flip/constants"
	"fmt"
	"slices"
	"sync"
	"sync/atomic"
	"time"
)

var sessionId atomic.Uint64

type WorkoutSession struct {
	Id            uint64
	WorkoutSlotId uint64
	Date          string
	SeatsBooked   map[uint64]bool
	SeatLock      sync.Mutex
	normalQueue   []*Booking
	premiumQueue  []*Booking
	Status        int
}

func NewWorkoutSession(slotId uint64) error {
	ws, err := GetWorkSlotById(slotId)
	if err != nil {
		return err
	}

	for i := 0; i < 7; i++ {
		duration := 24 * time.Duration(i) * time.Hour
		curDate := time.Now().Add(duration).Format("2006-01-02")
		if _, ok := WorkoutSchedule[curDate]; !ok {
			WorkoutSchedule[curDate] = make(map[int][]uint64)
		}
		if !slices.Contains(ws.SlotDays, constants.GetWeekday(curDate)) {
			continue
		}
		session := &WorkoutSession{
			Id:            sessionId.Add(1),
			WorkoutSlotId: slotId,
			SeatsBooked:   make(map[uint64]bool),
			SeatLock:      sync.Mutex{},
			Date:          curDate,
		}
		WorkoutSchedule[curDate][ws.WorkoutType] = append(WorkoutSchedule[curDate][ws.WorkoutType], session.Id)
		WorkoutSessions[session.Id] = session
	}
	return nil
}

func FillSessionsForWeek(centreId uint64) {
	centre := Centres[centreId]
	for _, slotIds := range centre.WorkoutTypesSchedule {
		for i := range slotIds {
			NewWorkoutSession(slotIds[i])
		}
	}
}

func GetWorkoutSessionById(id uint64) (*WorkoutSession, error) {
	_, ok := WorkoutSessions[id]
	if !ok {
		return nil, errors.New("invalid workout sessionId")
	}
	return WorkoutSessions[id], nil
}

func GetAvailableSessions(date string, workoutType int) []*WorkoutSession {
	sessionIds, ok := WorkoutSchedule[date][workoutType]
	if ok {
		var sessions []*WorkoutSession
		for i := range sessionIds {
			sessions = append(sessions, WorkoutSessions[sessionIds[i]])
		}
		return sessions
	}
	return nil
}

func (s *WorkoutSession) BookSlot(user *User) (uint64, error) {
	workoutSlot, _ := GetWorkSlotById(s.WorkoutSlotId)
	if workoutSlot.SlotType != constants.SLOT_TYPE_NORMAL && workoutSlot.SlotType != user.UserType {
		return 0, errors.New("this user is not eligible for this slot")
	}
	if err := user.CheckDateSlots(s.Date, workoutSlot.CentreId); err != nil {
		return 0, err
	}

	booking, _ := NewBooking(user.UserId, workoutSlot.CentreId, s.Id, s.Date)
	availableSeats := workoutSlot.NumberOfSeats - len(s.SeatsBooked)
	if availableSeats > 0 {
		s.SeatLock.Lock()
		defer s.SeatLock.Unlock()
		s.SeatsBooked[booking.Id] = true
		booking.UpdateBookingStatus(constants.Booking_Confirmed)
		return booking.Id, nil
	} else if user.UserType == constants.FK_VIP_USER {
		s.premiumQueue = append(s.premiumQueue, booking)
	} else {
		s.normalQueue = append(s.normalQueue, booking)
	}
	booking.UpdateBookingStatus(constants.Booking_Waiting)

	return 0, errors.New("slot fully booked")
}

func (s *WorkoutSession) GetNextBookingFromQueue() *Booking {
	var nextBooking *Booking
	if len(s.premiumQueue) > 0 {
		nextBooking = s.premiumQueue[0]
		s.premiumQueue = s.premiumQueue[1:]
	} else if len(s.normalQueue) > 0 {
		nextBooking = s.normalQueue[0]
		s.normalQueue = s.normalQueue[1:]
	}
	return nextBooking
}

func (s *WorkoutSession) CancelSession(booking *Booking) {
	s.SeatLock.Lock()
	defer s.SeatLock.Unlock()
	delete(s.SeatsBooked, booking.UserId)
	booking.UpdateBookingStatus(constants.Booking_Cancelled)

	workoutDetails, _ := json.Marshal(s)
	fmt.Println("Workout session cancelled by user", booking.UserId, string(workoutDetails))

	for nextBooking := s.GetNextBookingFromQueue(); nextBooking != nil; {
		user, _ := GetUserById(nextBooking.UserId)
		if user.CheckDateSlots(nextBooking.Date, nextBooking.CentreId) == nil {
			s.SeatsBooked[nextBooking.UserId] = true
			nextBooking.UpdateBookingStatus(constants.Booking_Confirmed)
			break
		}
	}
}
