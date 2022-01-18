package maxprofit

import "math"

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		minPrice = int(math.Min(float64(minPrice), float64(price)))
		maxProfit = int(math.Max(float64(maxProfit), float64(price-minPrice)))
	}

	return maxProfit
}
