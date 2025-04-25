package theatre

import "sync/atomic"

var seatNo atomic.Uint64

type Seat struct {
	SeatNo       uint64
	RowNo        int
	SeatCategory int
}

func NewSeat(rowNo, category int) *Seat {

	return &Seat{
		SeatNo:       seatNo.Add(1),
		RowNo:        rowNo,
		SeatCategory: category,
	}
}
