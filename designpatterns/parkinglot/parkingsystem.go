package parkinglot

type ParkingSystem struct {
	big, medium, small int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{
		big:    big,
		medium: medium,
		small:  small,
	}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big == 0 {
			return false
		} else {
			this.big--
			return true
		}
	case 2:
		if this.medium == 0 {
			return false
		} else {
			this.medium--
			return true
		}
	case 3:
		if this.small == 0 {
			return false
		} else {
			this.small--
			return true
		}
	default:
		return false
	}
}
