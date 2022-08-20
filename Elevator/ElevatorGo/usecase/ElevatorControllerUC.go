package usecase

import (
	"dotpe/projects/domain/entity"
	"math/rand"
)

type ElevatorControllerUC struct {
	StartFloor    int
	EndFloor      int
	NoofElevators int
	Elevators     map[int]entity.Elevator
	Floors        map[int]entity.Floor
}

func NewDispacherUcase(startfl, endfl, elevators int) ElevatorControllerUC {
	d := ElevatorControllerUC{
		StartFloor:    startfl,
		EndFloor:      endfl,
		NoofElevators: elevators,
		Elevators:     make(map[int]entity.Elevator),
		Floors:        make(map[int]entity.Floor),
	}
	for i := 0; i < d.NoofElevators; i++ {
		d.Elevators[i+1] = entity.NewElevator(rand.Intn(endfl-startfl)+startfl, startfl, endfl)
	}
	return d
}

func (d *ElevatorControllerUC) AddRequest(typeOfButton string, elevatorId, floorNo, dir int) {
	switch typeOfButton {
	case "ElevatorButton":
		d.Elevators[elevatorId].AddRequest(floorNo)
	case "FloorButton":
		d.CheckClosestElevator(floorNo, dir)
	}
}

func (d *ElevatorControllerUC) CheckClosestElevator(floorNo, dir int) int {
	minDist := d.EndFloor - d.StartFloor
	assignedElevator := d.EndFloor + 1
	for id, ele := range d.Elevators {
		curDiff := ele.CurrentFloor - floorNo
		switch ele.Direction {
		case entity.IDLE:
			if curDiff >= 0 && curDiff < minDist {
				minDist = curDiff
				assignedElevator = id
			} else if curDiff < 0 && (-curDiff) < minDist {
				minDist = -curDiff
				assignedElevator = id
			}
		case entity.UP:
			if curDiff >= 0 {
				minDist = curDiff
				assignedElevator = id
			} else if curDiff < 0 && (-curDiff) < minDist {
				minDist = -curDiff
				assignedElevator = id
			}

		}
	}
	return assignedElevator
}
