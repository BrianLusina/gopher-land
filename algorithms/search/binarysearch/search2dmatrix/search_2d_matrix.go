package search2dmatrix

func searchMatrix(matrix [][]int, target int) bool {
	nCols := len(matrix[0])

	for _, row := range matrix {
		left, right := 0, nCols-1
		firstTrueIndex := -1

		for left <= right {
			mid := (right + left) / 2
			if row[mid] >= target {
				firstTrueIndex = mid
				right = mid - 1
			} else if row[mid] < target {
				left = mid + 1
			}
		}

		if firstTrueIndex != -1 && row[firstTrueIndex] == target {
			return true
		}
	}
	return false
}
