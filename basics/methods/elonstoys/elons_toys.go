package elon

import "fmt"

// Drive drives the car forward draining its battery and covering distance
func (car *Car) Drive() {
	if car.batteryDrain > car.battery {
		return
	}

	car.battery -= car.batteryDrain
	car.distance += car.speed
}

// DisplayDistance displays the distance driven by the car
func (car *Car) DisplayDistance() string {
	return fmt.Sprintf("Driven %d meters", car.distance)
}

// DisplayBattery displays the current percent of the battert
func (car *Car) DisplayBattery() string {
	return fmt.Sprintf("Battery at %d%%", car.battery)
}

// CanFinish returns true if the car can finish the provided track distance
func (car *Car) CanFinish(trackDistance int) bool {
	maxDistance := car.battery / car.batteryDrain * car.speed
	return maxDistance >= trackDistance
}
