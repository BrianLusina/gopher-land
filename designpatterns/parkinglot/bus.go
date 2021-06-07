package parkinglot

type Bus struct {
	Vehicle
}

func NewBus(licensePlate string) Bus {
	return Bus{
		Vehicle: Vehicle{
			Size:         int(LARGE),
			LicensePlate: licensePlate,
			SpotSize:     5,
			SpotsTaken:   []ParkingSpot{},
		},
	}
}

func (this *Bus) CanFitInSpot(parkingSpot *ParkingSpot) bool {
	if this.SpotSize == parkingSpot.SpotSize {
		return true
	}
	return false
}
