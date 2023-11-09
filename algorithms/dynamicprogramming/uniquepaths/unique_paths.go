package uniquepaths

import (
	"fmt"
	mathUtils "gopherland/math/utils"
	"gopherland/pkg/utils"
)

// uniquePathsMath uses a math formula to determine the number of unique paths.
// We need to make n-1 + m-1 steps in total. How many ways to choose from m-1 right steps and n-1 down steps out of the total steps?
func uniquePathsMath(m int, n int) int {
	ans := 1
	x := m + n - 2
	k := mathUtils.Min(n-1, m-1)

	for i := 1; i < k+1; i++ {
		ans *= x
		x -= 1
		ans /= i
	}

	return ans
}

// uniquePathsTopDown uses top-down approach with memoization
// We are going to traverse all the unique paths, and store the values of the number of unique paths of each cell
// in our cache. Slight difference, we can start at m,n and traverse towards 0,0 to get the same result, which allows
// us to reuse the function as our recursive function.
//
// Complexity Analysis:
// - Time Complexity: O(m*n)
// - Space Complexity: O(m*n)
func uniquePathsTopDown(m int, n int) int {
	cache := map[string]int{}

	var uniquePathsHelper func(row, col int) int
	uniquePathsHelper = func(row, col int) int {
		search := fmt.Sprintf("%d#%d", row, col)

		if x, ok := cache[search]; ok {
			return x
		}

		var result int
		if row == 1 || col == 1 {
			result = 1
		} else {
			result = uniquePathsHelper(row-1, col) + uniquePathsHelper(row, col-1)
		}

		cache[search] = result
		return result
	}

	return uniquePathsHelper(m, n)
}

// uniquePathsBottomUp uses bottom-up approach
// Complexity Analysis:
// - Time Complexity: O(m*n)
// - Space Complexity: O(n)
func uniquePathsBottomUp(m int, n int) int {
	row := utils.MakeArray(n, 1)

	// go through all rows except the last one
	for i := 0; i < m-1; i++ {
		newRow := utils.MakeArray(n, 1)

		// go through every column except the right most column
		// because the last value in every row is 1
		// start at second to last position and
		// keep going until we get to the beginning (reverse order)

		for j := n - 2; j > -1; j-- {
			// right value + value below
			newRow[j] = newRow[j+1] + row[j]
		}
		// update the row
		row = newRow
	}

	return row[0]
}
