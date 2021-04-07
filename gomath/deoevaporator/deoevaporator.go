package deoevaporator

import "math"

func Evaporator(_ float64, evapPerDay int, threshold int) int {
	return int(math.Ceil(math.Log(float64(threshold)/100) / math.Log(1-float64(evapPerDay)/100)))
}
