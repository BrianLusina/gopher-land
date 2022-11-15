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
