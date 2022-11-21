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

// MaxProfitTwoPointers finds the maximum possible profit from a slice of prices. This is a Variation of MaxProfit using 2 pointers
// Complexity Analysis:
// Space: O(1), no extra memory is used
// Time: O(n), where n is the size of the input list & we iterate through the list only once
func MaxProfitTwoPointers(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	currentMaxProfit := 0

	left, right := 0, 1

	for right < len(prices) {
		low, high := prices[left], prices[right]
		if low < high {
			profit := high - low
			currentMaxProfit = int(math.Max(float64(currentMaxProfit), float64(profit)))
		} else {
			// We shift the left pointer to the right pointers position because at this point we know the right pointer
			// is at the lowest possible profit, shifting left by 1 will always give us negative profit. However,
			// shifting it to where the right pointer is will give us the next potential profit to find
			left = right
		}
		right++
	}

	return currentMaxProfit
}

// MaxProfit2 finds the total maximum possible profit from a slice of prices.
// Complexity Analysis:
// Space: O(1), no extra memory is used
// Time: O(n), where n is the size of the input list & we iterate through the list only once
func MaxProfit2(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	currentMaxProfit := 0

	left, right := 0, 1

	for right < len(prices) {
		previousPrice, nextPrice := prices[left], prices[right]
		if previousPrice < nextPrice {
			profit := nextPrice - previousPrice
			currentMaxProfit += profit
		}
		left++
		right++
	}

	return currentMaxProfit
}

func MaxProfit3(prices []int) int {
	n := len(prices)

	if n < 2 {
		return 0
	}
	forwardProfit, backwardProfit := make([]int, n), make([]int, n)
	earliestPrice := prices[0]

	for day := 1; day < n; day++ {
		earliestPrice = int(math.Min(float64(prices[day]), float64(earliestPrice)))
		forwardProfit[day] = int(math.Max(float64(forwardProfit[day-1]), float64(prices[day]-earliestPrice)))
	}

	latestPrice := prices[len(prices)-1]
	for day := n - 2; day > 0; day-- {
		latestPrice = int(math.Max(float64(latestPrice), float64(prices[day])))
		backwardProfit[day] = int(math.Max(float64(backwardProfit[day+1]), float64(latestPrice-prices[day])))
	}

	profit := 0
	for day := range prices {
		profit = int(math.Max(float64(profit), float64(forwardProfit[day]+backwardProfit[day])))
	}

	return profit
}
