package parkinglot

type MotorCycle struct {
	Vehicle
}

func NewMotorCycle(licensePlate string) MotorCycle {
	return MotorCycle{
		Vehicle: Vehicle{
			Size:         int(MOTORCYCLE),
			LicensePlate: licensePlate,
			SpotSize:     1,
			SpotsTaken:   []ParkingSpot{},
		},
	}
}

func (this *MotorCycle) CanFitInSpot(parkingSpot *ParkingSpot) bool {
	return true
}
