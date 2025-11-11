package settozero

import (
	"slices"
)

/*
*
setMaxZeroes sets all elements in a matrix to zero if any row or column has at least one zero element.

## Complexity Analysis

This solution is efficient in both time and space. The time complexity is O(rows × cols) because we scan through
the matrix a constant number of times. The space complexity is O(1) because we only use a couple of boolean
variables regardless of the matrix size, and we use the matrix itself for tracking.

Args:

	mat (List): input matrix

Returns:

	List[List[int]]: modified matrix with all elements in any row or column containing a zero element set to zero

*
*/
func setMatrixZeros(mat [][]int) [][]int {
	if len(mat) == 0 || len(mat[0]) == 0 {
		return mat
	}

	rowLength, colLength := len(mat), len(mat[0])

	// Step 1: Check if the first row and first column originally contain zeros
	// We need this info because we'll use them as markers
	firstRowHasZero := slices.ContainsFunc(mat[0], func(e int) bool { return e == 0 })
	var firstColHasZero = false
	for r := range rowLength {
		if mat[r][0] == 0 {
			firstColHasZero = true
			break
		}
	}

	// Step 2: Use first row and column as markers for other rows/columns
	// Scan the interior of the matrix (excluding first row and column)
	for r := 1; r < rowLength; r++ {
		for c := 1; c < colLength; c++ {
			if mat[r][c] == 0 {
				// Mark this row and column by setting the corresponding
				// positions in the first row and column to zero
				mat[r][0] = 0 // mark row as 0
				mat[0][c] = 0 // mark column as 0
			}
		}
	}

	// Step 3: Use the markers to set zeros in the interior of the matrix
	// We process from [1][1] onwards to avoid corrupting our markers
	for r := 1; r < rowLength; r++ {
		for c := 1; c < colLength; c++ {
			// If this row or column was marked, set this cell to zero
			if mat[r][0] == 0 || mat[0][c] == 0 {
				mat[r][c] = 0
			}
		}
	}

	// Step 4: Finally, handle the first row and column based on our original checks
	if firstRowHasZero {
		for c := range colLength {
			mat[0][c] = 0
		}
	}
	if firstColHasZero {
		for r := range rowLength {
			mat[r][0] = 0
		}
	}

	return mat
}
