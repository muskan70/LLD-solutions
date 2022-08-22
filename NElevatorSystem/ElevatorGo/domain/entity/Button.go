package entity

type FloorButton struct {
	Direction bool
}

type ElevatorButton struct {
	FloorNo int
}

func NewFloorButton(dir bool) FloorButton {
	return FloorButton{Direction: dir}
}

func NewElevatorButton(floorNo int) ElevatorButton {
	return ElevatorButton{FloorNo: floorNo}
}

func (e *ElevatorButton) Press() {

}
