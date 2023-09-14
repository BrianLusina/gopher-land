package threepointers

import (
	"gopherland/math/utils"
	"math"
)

// minimize minimizes the maximum difference
// Windowing strategy works here.
// Take 3 pointers i, j and k
//
// Initialize them to point to the start of arrays a, b and c
// Find min of i, j and k.
// Compute max(i, j, k) - min(i, j, k).
// If the new result is less than the current result, change it to the new result.
// Increment the pointer of the array which contains the minimum.
//
// Note that we increment the pointer of the array which has the minimum, because our goal is to decrease the difference.
// Increasing the maximum pointer is definitely going to increase the difference. Increasing the second maximum pointer
// can potentially increase the difference (however, it certainly will not decrease the difference).
func minimize(a, b, c []int) int {
	currentMin := math.MaxInt64
	i, j, k := 0, 0, 0

	for i < len(a) && j < len(b) && k < len(c) {
		aNum, bNum, cNum := a[i], b[j], c[k]

		ab := utils.Abs(aNum - bNum)
		bc := utils.Abs(bNum - cNum)
		ca := utils.Abs(cNum - aNum)

		current := utils.Max(utils.Max(ab, bc), ca)
		currentMin = utils.Min(current, currentMin)

		minOfThree := utils.Min(utils.Min(aNum, bNum), cNum)

		if aNum == minOfThree {
			i++
		} else if bNum == minOfThree {
			j++
		} else {
			k++
		}
	}

	return currentMin
}
