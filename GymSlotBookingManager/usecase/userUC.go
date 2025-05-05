package usecase

import (
	"errors"
	"flip/constants"
	"flip/model"
	"sort"
	"time"
)

type AvailableSessionsWrtUser struct {
	Distance float64
	CentreId uint64
	Session  *model.WorkoutSession
}

func CreateWorkoutSlot(centreId uint64, workoutType int, slotTime, availableSeats, slotType int, days []int, charges float64) (uint64, error) {
	slotId, err := model.NewWorkoutSlot(centreId, workoutType, slotTime, availableSeats, slotType, days, charges)
	if err != nil {
		return 0, err
	}

	centre, _ := model.GetCentre(centreId)

	if err := centre.AddWorkoutSlot(slotId); err != nil {
		return 0, err
	}
	return slotId, nil
}

func GetAvailableSessions(userId uint64, workoutType int, date string) ([]*AvailableSessionsWrtUser, error) {
	user, _ := model.GetUserById(userId)
	var availableSessions []*AvailableSessionsWrtUser
	slots := model.GetAvailableSessions(date, workoutType)

	for x := range slots {
		slotDetails, _ := model.GetWorkSlotById(slots[x].WorkoutSlotId)
		if (slotDetails.SlotType == constants.SLOT_TYPE_NORMAL || slotDetails.SlotType == user.UserType) && (slotDetails.NumberOfSeats-len(slots[x].SeatsBooked) > 0) {
			centre, _ := model.GetCentre(slotDetails.CentreId)
			availableSessions = append(availableSessions, &AvailableSessionsWrtUser{
				Distance: model.GetDistance(centre.Location, user.Location),
				Session:  slots[x],
				CentreId: slotDetails.CentreId,
			})
		}
	}
	if len(availableSessions) == 0 {
		return nil, errors.New("no sessions available for today")
	}
	sort.Slice(availableSessions, func(i, j int) bool {
		return availableSessions[i].Distance < availableSessions[j].Distance
	})
	return availableSessions, nil
}

func BookSession(userId uint64, session *model.WorkoutSession) (uint64, error) {
	user, err := model.GetUserById(userId)
	if err != nil {
		return 0, err
	}

	slotDetails, err := model.GetWorkSlotById(session.WorkoutSlotId)
	if err != nil {
		return 0, errors.New("invalid centre")
	}

	if err := user.CheckDateSlots(session.Date, slotDetails.CentreId); err != nil {
		return 0, err
	}

	bookingId, err := session.BookSlot(user)
	if err != nil {
		return 0, err
	}
	return bookingId, nil
}

func CancelNextSession(userId uint64) error {
	user, _ := model.GetUserById(userId)
	curDate := time.Now().Format("2006-01-02")
	bookings := user.GetAllConfirmedBookingsByDate(curDate)
	if len(bookings) == 0 {
		return errors.New("no confirmed bookings for today")
	}
	session, _ := model.GetWorkoutSessionById(bookings[0].SessionId)
	session.CancelSession(bookings[0])
	return nil
}
