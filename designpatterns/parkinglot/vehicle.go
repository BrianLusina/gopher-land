package parkinglot

type attributes interface {
	CanFitInSpot(parkingSpot *ParkingSpot) bool
}

type Vehicle struct {
	Size         int
	LicensePlate string
	SpotSize     int
	SpotsTaken   []ParkingSpot
	Attributes   attributes
}

func (this *Vehicle) ClearSpots() {
	for _, spot := range this.SpotsTaken {
		spot.RemoveVehicle(this)
	}
}

func (this *Vehicle) TakeSpot(parkingSpot ParkingSpot) {
	this.SpotsTaken = append(this.SpotsTaken, parkingSpot)
}

func (this *Vehicle) CanFitInSpot(parkingSpot ParkingSpot) bool {
	return this.Size > parkingSpot.SpotSize
}
