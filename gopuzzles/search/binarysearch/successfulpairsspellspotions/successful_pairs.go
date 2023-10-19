package successfulpairsspellspotions

import "sort"

// successfulPairs finds the successful pairs of spells and potions that meet the given success criteria
// Space: O(n) - storing the result
// Time: O(mlogm + nlogm), where O(mlogm) is sorting the potions and O(nlogm) is performing the binary search on the potions
func successfulPairs(spells []int, potions []int, success int64) []int {
	var validPosition func([]int, int64, int) int
	validPosition = func(sortedPotions []int, target int64, currentSpell int) int {
		potionNeeded := (target + int64(currentSpell) - 1) / int64(currentSpell)
		left, right := 0, len(sortedPotions)

		for left < right {
			middle := (left + right) >> 1

			if sortedPotions[middle] >= int(potionNeeded) {
				right = middle
			} else {
				left = middle + 1
			}
		}
		return left
	}

	sort.IntSlice(potions).Sort()

	result := []int{}
	for _, spell := range spells {
		count := len(potions) - validPosition(potions, success, spell)
		result = append(result, count)
	}

	return result
}
