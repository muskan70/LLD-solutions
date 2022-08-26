package main

import (
	"fmt"
	"parking/constants"
	"parking/parkingLot"

	"parking/ticket"
)

func main() {
	pl := parkingLot.NewParkingLot(1, constants.ParkingStrategyTight)
	pl.InitializeSpace(1, constants.VehicleTypeSmallCar, 10)

	pc := constants.InitializeParkingCharges()

	tk1 := ticket.CreateTicket(pl, 2346, constants.VehicleTypeSmallCar, 5)
	fmt.Println(tk1.TicketNo, tk1.FloorNo, tk1.SlotNo)
	tk2 := ticket.CreateTicket(pl, 236, constants.VehicleTypeSmallCar, 10)
	fmt.Println(tk2.TicketNo, tk2.FloorNo, tk2.SlotNo)
	tk3 := ticket.CreateTicket(pl, 234, constants.VehicleTypeSmallCar, 60)
	fmt.Println(tk3.TicketNo, tk3.FloorNo, tk3.SlotNo)
	if tk2 != nil {
		fmt.Println("ticket defined")
		tk2.GetTicketCharges(pl, 305, pc)
		fmt.Println(tk2.TotalCharges)
	}
	fmt.Println(pl.CheckAvailableSlot(constants.VehicleTypeSmallCar))
}
