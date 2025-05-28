package business

import (
	"sliceMC/constants"
	"sliceMC/entity"
)

type HotelManager struct {
	Rooms map[uint64]*entity.Room
}

func NewHotelManager() *HotelManager {
	return &HotelManager{
		Rooms: make(map[uint64]*entity.Room),
	}
}

func (hm *HotelManager) CreateRoom(roomType int) uint64 {
	room := entity.NewRoom(roomType)
	hm.Rooms[room.Id] = room
	return room.Id
}

func (hm *HotelManager) GetAllRoomsByStatus(roomStatus int) []*entity.Room {
	var rooms []*entity.Room
	for _, room := range hm.Rooms {
		if room.BookStatus == roomStatus {
			rooms = append(rooms, room)
		}
	}
	return rooms
}

func (hm *HotelManager) CheckRoomAvailability(roomId uint64) bool {
	return hm.Rooms[roomId].BookStatus == constants.ROOM_STATUS_AVAILABLE
}

func (hm *HotelManager) BookRoom(roomId uint64) error {
	return hm.Rooms[roomId].BookRoom()
}
