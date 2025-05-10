package services

import (
	"errors"
	"sync"
	"ticketBooking/constants"
	"ticketBooking/models"
)

type BookingService struct {
	Bookings map[uint64]*models.Booking
}

func NewBookingService() *BookingService {
	return &BookingService{
		Bookings: make(map[uint64]*models.Booking),
	}
}

func (b *BookingService) GetBookingById(bookingId uint64) *models.Booking {
	return b.Bookings[bookingId]
}

func (b *BookingService) BookTickets(userId, showId uint64, seatIds []string) error {
	show := showManager.GetShowById(showId)
	for _, id := range seatIds {
		if !show.CheckSeatAvailability(id) {
			return errors.New("seats not available")
		}
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(seatIds))
	for _, id := range seatIds {
		go show.UpdateSeatStatus(id, constants.SEAT_STATUS_BOOKED, wg)
	}
	wg.Wait()

	ticket := models.NewBooking(showId, userId, seatIds)
	b.Bookings[ticket.Id] = ticket
	return nil
}

func (b *BookingService) CancelTickets(bookingId uint64) {
	booking := b.GetBookingById(bookingId)
	show := showManager.GetShowById(booking.ShowId)

	wg := &sync.WaitGroup{}
	wg.Add(len(booking.SeatIds))
	for _, id := range booking.SeatIds {
		go show.UpdateSeatStatus(id, constants.SEAT_STATUS_AVAILABLE, wg)
	}
	wg.Wait()

	booking.CancelBooking()
}
