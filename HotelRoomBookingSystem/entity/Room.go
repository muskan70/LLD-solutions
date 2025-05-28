package entity

import (
	"errors"
	"sliceMC/constants"
	"sync/atomic"
)

var roomId atomic.Uint64

type Room struct {
	Id         uint64
	BookStatus int
	RoomType   int
}

func NewRoom(roomType int) *Room {
	return &Room{
		Id:         roomId.Add(1),
		BookStatus: constants.ROOM_STATUS_AVAILABLE,
		RoomType:   roomType,
	}
}

func (r *Room) BookRoom() error {
	if r.BookStatus == constants.ROOM_STATUS_AVAILABLE {
		r.BookStatus = constants.ROOM_STATUS_HELD
		return nil
	}
	return errors.New("room not free")
}

func (r *Room) ConfirmBooking() error {
	if r.BookStatus == constants.ROOM_STATUS_HELD {
		r.BookStatus = constants.ROOM_STATUS_BOOKED
		return nil
	}
	return errors.New("room status not held")
}

func (r *Room) Free() error {
	if r.BookStatus == constants.ROOM_STATUS_AVAILABLE {
		return errors.New("room already available")
	}
	r.BookStatus = constants.ROOM_STATUS_AVAILABLE
	return nil
}
