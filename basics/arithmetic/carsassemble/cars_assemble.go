package carsassemble

const costToProduceTen = 95000
const costToProduceOne = 10000

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	percentage := successRate / 100
	return float64(productionRate) * percentage
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successfulCarsPerHour := CalculateWorkingCarsPerHour(productionRate, successRate)
	return int(successfulCarsPerHour / 60)
}

// CalculateCost works out the cost of producing the given number of cars
func CalculateCost(carsCount int) uint {
	groups := carsCount / 10
	rem := carsCount % 10
	cost := groups*costToProduceTen + rem*costToProduceOne
	return uint(cost)
}
