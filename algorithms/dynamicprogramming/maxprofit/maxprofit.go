package maxprofit

import (
	"gopherland/math/utils"
)

func MaxProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		minPrice = utils.Min(minPrice, price)
		maxProfit = utils.Max(maxProfit, price-minPrice)
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
			currentMaxProfit = utils.Max(currentMaxProfit, profit)
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
		earliestPrice = utils.Min(prices[day], earliestPrice)
		forwardProfit[day] = utils.Max(forwardProfit[day-1], prices[day]-earliestPrice)
	}

	latestPrice := prices[len(prices)-1]
	for day := n - 2; day > 0; day-- {
		latestPrice = utils.Max(latestPrice, prices[day])
		backwardProfit[day] = utils.Max(backwardProfit[day+1], latestPrice-prices[day])
	}

	profit := 0
	for day := range prices {
		profit = utils.Max(profit, forwardProfit[day]+backwardProfit[day])
	}

	return profit
}

func MaxProfitWithFee(prices []int, fee int) int {
	if len(prices) == 0 {
		return 0
	}

	cashWithShares := -prices[0]
	cashWithoutShares := 0

	for i := 1; i < len(prices); i++ {
		// Maximum cash in hand with shares
		// Either
		// 1. withold prev share in which case your cash in hand will not change,
		// 2. or assume there was no currently bought share but you want to buy it today -
		//         In this case: your current cash in hand with shares will be your previous cash
		//         in hand without shares minus buying price of the share today.
		cashWithShares = utils.Max(cashWithShares, cashWithoutShares-prices[i])

		// Maximum cash in hand without shares
		// Either
		// 1. withold money without shares in which case your cash in hand will not change,
		// 2. or assume you previously bought the share and you are going to sell that today -
		//         In this case : your current cash in hand without shares will be your previous cash
		//         in hand with shares plus the current selling price minus transaction fee
		cashWithoutShares = utils.Max(cashWithoutShares, cashWithShares+prices[i]-fee)
	}

	return cashWithoutShares
}
