package tanktruck

import "math"

func TankVol(h, d, vt int) int {
	radius := d / 2
	cylinderHeight := vt / int((math.Pi * math.Pow(float64(radius), float64(2))))
	triangleHeight := radius - h
	theta := math.Acos(float64(triangleHeight / radius))

	// or base = radius * sin(theta)
	base := radius * int(math.Sin(theta))

	triangleArea := (base * triangleHeight) / 2
	sectorArea := (radius * radius * int(theta)) / 2

	remainderArea := (sectorArea - triangleArea) * 2
	return cylinderHeight * remainderArea
}
