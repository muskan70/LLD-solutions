package entity

import "github.com/emirpasic/gods/trees/binaryheap"

const (
	UP   = 1
	DOWN = -1
	IDLE = 0
)

var x = 1

type Elevator struct {
	ElevatorId   int
	Direction    int
	CurrentFloor int
	StartFloor   int
	EndFloor     int
	Door         bool
	ButtonPanel  map[int]ElevatorButton
	UpRequests   *binaryheap.Heap
	DownRequests *binaryheap.Heap
}

func NewElevator(cur, startFl, endFl int) Elevator {
	e := Elevator{
		CurrentFloor: cur,
		StartFloor:   startFl,
		EndFloor:     endFl,
		Door:         false,
		Direction:    IDLE,
		ButtonPanel:  make(map[int]ElevatorButton),
		ElevatorId:   x,
	}
	for i := startFl; i <= endFl; i++ {
		e.ButtonPanel[i] = ElevatorButton{FloorNo: i}
	}
	x++
	return e
}

func (e *Elevator) CloseDoor() {
	e.Door = false
}

func (e *Elevator) OpenDoor() {
	e.Door = true
}
