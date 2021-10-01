package lasagna

const OvenTime = 40

func RemainingOvenTime(actual int) int {
	return OvenTime - actual
}

func PreparationTime(layers int) int {
	return layers * 2
}

func ElapsedTime(layers, timeInOven int) int {
	return PreparationTime(layers) + timeInOven
}
