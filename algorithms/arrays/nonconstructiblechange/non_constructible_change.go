package nonconstructiblechange

import (
	"gopherland/algorithms/sorting/mergesort"
	"sort"
)

// nonConstructibleChange returns the smallest change that cannot be made using the given coins
func nonConstructibleChange(coins []int) int {
	// sort the coins and return a copy of the sort
	// sort.Ints(coins)
	sortedCoins := mergesort.MergeSort(coins)
	currentChange := 0
	for _, coin := range sortedCoins {
		if coin > currentChange+1 {
			return currentChange + 1
		}
		currentChange += coin
	}
	return currentChange + 1
}

// nonConstructibleChangeV2 returns the smallest change that cannot be made using the given coins, sorts the coins in place
func nonConstructibleChangeV2(coins []int) int {
	// sort the coins and return a copy of the sort
	sort.Ints(coins)
	currentChange := 0
	for _, coin := range coins {
		if coin > currentChange+1 {
			return currentChange + 1
		}
		currentChange += coin
	}
	return currentChange + 1
}
