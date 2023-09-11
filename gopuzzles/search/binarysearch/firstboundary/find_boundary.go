package firstboundary

func findBoundary(arr []bool) int {
	left := 0
	right := len(arr) - 1
	boundaryIndex := -1

	for left <= right {
		mid := (left + right) / 2
		if arr[mid] {
			boundaryIndex = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return boundaryIndex
}
