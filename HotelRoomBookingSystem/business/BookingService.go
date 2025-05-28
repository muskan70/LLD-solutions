package business

import (
	"errors"
	"sliceMC/constants"
	"sliceMC/entity"
	"time"
)

type BookingService struct {
	HotelManager    *HotelManager
	Bookings        map[uint64]*entity.Booking
	UserHeldBooking map[uint64]uint64
}

func NewBookingService(hm *HotelManager) *BookingService {
	return &BookingService{
		HotelManager:    hm,
		Bookings:        make(map[uint64]*entity.Booking),
		UserHeldBooking: make(map[uint64]uint64),
	}
}

func (bs *BookingService) CreateBooking(userId, roomId uint64) (uint64, error) {
	if _, ok := bs.UserHeldBooking[userId]; ok {
		return 0, errors.New("user has already holding a room for booking")
	}
	booking := entity.NewBooking(userId, roomId)
	bs.Bookings[booking.Id] = booking
	if err := bs.HotelManager.Rooms[roomId].BookRoom(); err != nil {
		bs.Bookings[booking.Id].Status = constants.BOOKING_STATUS_CANCELLED
		return 0, err
	} else {
		bs.Bookings[booking.Id].Status = constants.BOOKING_STATUS_PENDING
		bs.UserHeldBooking[userId] = booking.Id
		return booking.Id, nil
	}
}

func (bs *BookingService) ConfirmBooking(bookingId uint64) error {
	booking := bs.Bookings[bookingId]
	if _, ok := bs.UserHeldBooking[booking.UserId]; !ok {
		return errors.New("this booking Id is not eligible for confirming")
	}
	err := bs.HotelManager.Rooms[booking.RoomId].ConfirmBooking()
	if err != nil {
		bs.Bookings[bookingId].Status = constants.BOOKING_STATUS_CANCELLED
	} else {
		bs.Bookings[bookingId].Status = constants.BOOKING_STATUS_CONFIRMED
	}
	delete(bs.UserHeldBooking, booking.UserId)
	return err
}

func (bs *BookingService) CancelBooking(bookingId uint64) error {
	booking := bs.Bookings[bookingId]
	if booking.Status == constants.BOOKING_STATUS_CANCELLED {
		return errors.New("this booking Id is already cancelled")
	}
	delete(bs.UserHeldBooking, booking.UserId)
	if err := bs.HotelManager.Rooms[booking.RoomId].Free(); err != nil {
		return err
	} else {
		bs.Bookings[bookingId].Status = constants.BOOKING_STATUS_CANCELLED
		return nil
	}
}

func (bs *BookingService) CheckExpiredBookings() {
	for _, bookingId := range bs.UserHeldBooking {
		booking := bs.Bookings[bookingId]
		if booking.Timestamp.Add(10 * time.Minute).Before(time.Now()) {
			delete(bs.UserHeldBooking, booking.UserId)
			bs.HotelManager.Rooms[booking.RoomId].Free()
			bs.Bookings[bookingId].Status = constants.BOOKING_STATUS_CANCELLED
		}
	}
}
