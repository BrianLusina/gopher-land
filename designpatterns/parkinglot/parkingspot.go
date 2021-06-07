package parkinglot

type ParkingSpot struct {
	Level       int
	Row         int
	SpotNumber  int
	SpotSize    int
	VehicleSize int
	Vehicle     *Vehicle
}

func NewParkingSpot(level, row, spotNumber, spotSize, vehicleSize int) ParkingSpot {
	return ParkingSpot{
		Level:       level,
		Row:         row,
		SpotNumber:  spotNumber,
		SpotSize:    spotSize,
		VehicleSize: vehicleSize,
		Vehicle:     nil,
	}
}

func (this *ParkingSpot) IsAvailable() bool {
	if this.Vehicle == nil {
		return true
	}
	return false
}

func (this *ParkingSpot) CanFitVehicle(vehicle Vehicle) bool {
	if this.Vehicle != nil {
		return false
	}
	return vehicle.Attributes.CanFitInSpot(this)
}

func (this *ParkingSpot) ParkVehicle(vehicle Vehicle) {
	this.SpotSize -= vehicle.SpotSize
}

func (this *ParkingSpot) RemoveVehicle(vehicle *Vehicle) {
	this.SpotSize += vehicle.SpotSize
}
