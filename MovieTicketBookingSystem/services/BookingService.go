package services

import (
	"errors"
	"fmt"
	"sync"
	"ticketBooking/constants"
	"ticketBooking/manager"
	"ticketBooking/models"
)

type BookingService struct {
	ShowManager      *manager.ShowManager
	Bookings         map[uint64]*models.Booking
	BookingsByUserId map[uint64][]uint64
}

func NewBookingService(showManager *manager.ShowManager) *BookingService {
	return &BookingService{
		Bookings:         make(map[uint64]*models.Booking),
		ShowManager:      showManager,
		BookingsByUserId: make(map[uint64][]uint64),
	}
}

func (b *BookingService) GetBookingById(bookingId uint64) *models.Booking {
	return b.Bookings[bookingId]
}

func (b *BookingService) GetBookingByUserId(id uint64) []uint64 {
	return b.BookingsByUserId[id]
}

func (b *BookingService) BookTickets(userId, showId uint64, seatIds []string) (uint64, error) {
	show := b.ShowManager.GetShowById(showId)
	for _, id := range seatIds {
		if !show.CheckSeatAvailability(id) {
			return 0, errors.New("seats not available")
		}
	}
	wg := &sync.WaitGroup{}
	wg.Add(len(seatIds))
	for i := range seatIds {
		go show.UpdateSeatStatus(seatIds[i], constants.SEAT_STATUS_BOOKED, wg)
	}
	wg.Wait()

	ticket := models.NewBooking(showId, userId, seatIds)
	b.Bookings[ticket.Id] = ticket
	b.BookingsByUserId[userId] = append(b.BookingsByUserId[userId], ticket.Id)
	return ticket.Id, nil
}

func (b *BookingService) CancelTickets(bookingId uint64) {
	booking := b.GetBookingById(bookingId)
	show := b.ShowManager.GetShowById(booking.ShowId)

	wg := &sync.WaitGroup{}
	wg.Add(len(booking.SeatIds))
	for _, id := range booking.SeatIds {
		go show.UpdateSeatStatus(id, constants.SEAT_STATUS_AVAILABLE, wg)
	}
	wg.Wait()

	booking.CancelBooking()
	fmt.Println("bookingId cancelled:", bookingId)
}
