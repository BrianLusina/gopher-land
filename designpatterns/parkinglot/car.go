package parkinglot

type Car struct {
	Vehicle
}

func NewCar(licensePlate string) Car {
	return Car{
		Vehicle: Vehicle{
			Size:         int(COMPACT),
			LicensePlate: licensePlate,
			SpotSize:     1,
			SpotsTaken:   []ParkingSpot{},
		},
	}
}

func (this *Car) CanFitInSpot(parkingSpot *ParkingSpot) bool {
	if this.SpotSize == int(LARGE) || this.SpotSize == int(COMPACT) {
		return true
	}
	return false
}
