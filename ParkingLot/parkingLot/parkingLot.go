package parkingLot

import "parking/floor"

type ParkingLot struct {
	NoofFloors      int
	ParkingStrategy int
	Floors          map[int]*floor.Floor
}

func NewParkingLot(floors int, parkingStrategy int) *ParkingLot {
	return &ParkingLot{NoofFloors: floors, ParkingStrategy: parkingStrategy, Floors: make(map[int]*floor.Floor)}
}

func (p *ParkingLot) InitializeSpace(floorNo, typeOfVehicle, spaces int) {
	if f, ok := p.Floors[floorNo]; ok {
		f.AddSpace(typeOfVehicle, spaces)
		p.Floors[floorNo] = f
	} else {
		f = floor.NewFloor(floorNo)
		f.AddSpace(typeOfVehicle, spaces)
		p.Floors[floorNo] = f
	}
}

func (p *ParkingLot) CheckAvailableSlot(typeofVehicle int) map[int]int {
	slots := make(map[int]int)
	for i, f := range p.Floors {
		if sp := f.CheckAvailability(typeofVehicle); sp > 0 {
			slots[i] = sp
		}
	}
	return slots
}
