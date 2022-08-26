package ticket

import (
	"fmt"
	"parking/constants"

	"parking/parkingLot"
)

type Ticket struct {
	TicketNo      int
	VehicleNo     int
	TypeofVehicle int
	TimeOfEntry   int
	TimeOfExit    int
	TimeTaken     int
	TotalCharges  float64
	FloorNo       int
	SlotNo        int
}

var x = 0

func CreateTicket(pl *parkingLot.ParkingLot, vehicleNo, typeofVehicle int, timeOfEntry int) *Ticket {
	slots := pl.CheckAvailableSlot(typeofVehicle)
	if len(slots) == 0 {
		return nil
	}
	floorNo := 0
	for key := range slots {
		floorNo = key
		break
	}
	fmt.Println(slots)
	slotNo := pl.Floors[floorNo].AllocateSpace(typeofVehicle)
	x++
	fmt.Println(pl.CheckAvailableSlot(typeofVehicle))
	return &Ticket{TicketNo: x, VehicleNo: vehicleNo, TypeofVehicle: typeofVehicle, TimeOfEntry: timeOfEntry, FloorNo: floorNo, SlotNo: slotNo}
}

func (t *Ticket) GetTicketCharges(pl *parkingLot.ParkingLot, timeOfExit int, pc *constants.ParkingCharges) {
	t.TimeOfExit = timeOfExit
	t.TimeTaken = timeOfExit - t.TimeOfEntry
	t.TotalCharges = pc.GetCharges(t.TimeTaken, t.TypeofVehicle)
	pl.Floors[t.FloorNo].DeallocateSpace(t.TypeofVehicle, t.SlotNo)
}
