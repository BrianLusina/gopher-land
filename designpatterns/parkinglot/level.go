package parkinglot

type Level struct {
	Floor          int
	NumSpots       int
	AvailableSpots int
	ParkingSpots   []ParkingSpot
}

func NewLevel(floor, numSpots, availableSpots int, parkingSpots []ParkingSpot) Level {
	return Level{
		Floor:          floor,
		NumSpots:       numSpots,
		AvailableSpots: availableSpots,
		ParkingSpots:   parkingSpots,
	}
}

func (this *Level) parkVehicle(vehicle Vehicle) bool {
	spot, err := this.findAvailableSpot(vehicle)
	if err != nil {
		return false
	} else {
		spot.ParkVehicle(vehicle)
		return true
	}
}

func (this *Level) findAvailableSpot(vehicle Vehicle) (ParkingSpot, error) {
	return ParkingSpot{}, nil
}

func (this *Level) parkStartingAtSpot(parkingSpot ParkingSpot, vehicle Vehicle) {

}
