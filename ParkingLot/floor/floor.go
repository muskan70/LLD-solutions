package floor

import "fmt"

type Floor struct {
	floorNo        int
	spaces         map[int]int
	availableSpace map[int]int
	Slots          map[int][]Slot
	TotalSlots     int
}

type Slot struct {
	SlotNo    int
	Available bool
}

func NewFloor(floorNo int) *Floor {
	return &Floor{floorNo: floorNo, spaces: make(map[int]int), availableSpace: make(map[int]int), Slots: make(map[int][]Slot), TotalSlots: 0}
}

func (f *Floor) AddSpace(typeOfVehicle, spaces int) {
	f.spaces[typeOfVehicle] = spaces
	f.availableSpace[typeOfVehicle] = spaces
	slots := []Slot{}
	for i := 1; i <= spaces; i++ {
		slots = append(slots, Slot{SlotNo: i, Available: true})
	}
	f.Slots[typeOfVehicle] = slots
}

func (f *Floor) AllocateSpace(typeofVehicle int) int {
	f.availableSpace[typeofVehicle] = f.availableSpace[typeofVehicle] - 1
	slots := f.Slots[typeofVehicle]
	for i, slot := range slots {
		if slot.Available {
			slots[i].Available = false
			f.Slots[typeofVehicle] = slots
			return slot.SlotNo
		}
	}
	return 0
}

func (f *Floor) DeallocateSpace(typeofVehicle int, slotNo int) {
	f.availableSpace[typeofVehicle] = f.availableSpace[typeofVehicle] + 1
	slots := f.Slots[typeofVehicle]
	slots[slotNo-1].Available = true
	f.Slots[typeofVehicle] = slots
}

func (f *Floor) CheckAvailability(typeofVehicle int) int {
	for _, slot := range f.Slots[typeofVehicle] {
		fmt.Println(slot.SlotNo, slot.Available)
	}
	return f.availableSpace[typeofVehicle]
}
