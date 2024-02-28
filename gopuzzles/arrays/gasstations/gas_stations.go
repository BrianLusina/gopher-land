package gasstations

import "gopherland/math/utils"

// canCompleteCircuit If sum of gas is less than sum of cost, then there is no way to get through all stations.
// So while we loop through the stations we sum up, so that at the end we can check the sum.
// Otherwise, there must be one unique solution, so the first one I find is the right one.
// If the tank becomes negative, we restart because that can't happen.
// Complexity: Time Complexity is O(n) where n is the number of gas stations to go though
// Space Complexity is O(1) as no extra space is required
func canCompleteCircuit(gas []int, cost []int) int {
	if (utils.Sum(gas) - utils.Sum(cost)) < 0 {
		return -1
	}

	gasTank, startIndex := 0, 0

	for idx, gasStation := range gas {
		gasTank += gasStation - cost[idx]

		if gasTank < 0 {
			gasTank = 0
			startIndex = idx + 1
		}
	}

	return startIndex
}
