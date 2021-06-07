package parkinglot

type ParkingLot struct {
	NumLevels int
	Levels    []Level
}

func NewParkingLot(numLevels int, levels []Level) ParkingLot {
	return ParkingLot{
		NumLevels: numLevels,
		Levels:    levels,
	}
}

func (this *ParkingLot) ParkVehicle(vehicle Vehicle) bool {
	for _, level := range this.Levels {
		if level.parkVehicle(vehicle) {
			return true
		}
	}
	return false
}
