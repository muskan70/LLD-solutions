package constants

type ParkingCharges struct {
	charges map[int][]float64
}

func InitializeParkingCharges() *ParkingCharges {
	pc := ParkingCharges{charges: make(map[int][]float64)}
	pc.charges[VehicleType2wheeler] = []float64{10, 20, 30, 40, 50}
	pc.charges[VehicleType3wheeler] = []float64{10, 20, 30, 40, 50}
	pc.charges[VehicleTypeSmallCar] = []float64{20, 40, 60, 80, 100}
	pc.charges[VehicleTypeBigCar] = []float64{10, 20, 30, 40, 50}
	pc.charges[VehicleTypeTruck] = []float64{10, 20, 30, 40, 50}
	return &pc
}

func (ch *ParkingCharges) GetCharges(timeTaken int, typeOfVehicle int) float64 {
	timeIdx := 0
	if timeTaken <= 60 {
		timeIdx = 0
	} else if timeTaken > 60 && timeTaken <= 120 {
		timeIdx = 1
	} else if timeTaken > 120 && timeTaken <= 300 {
		timeIdx = 2
	} else if timeTaken > 300 && timeTaken <= 450 {
		timeIdx = 3
	} else if timeTaken > 450 {
		timeIdx = 4
	}
	return ch.charges[typeOfVehicle][timeIdx]
}
