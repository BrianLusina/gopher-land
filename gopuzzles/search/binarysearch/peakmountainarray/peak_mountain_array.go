package peakmountainarray

func peakIndexInMountainArray(arr []int) int {
	left := 0
	right := len(arr) - 1
	k := -1

	for left <= right {
		mid := (left + right) / 2

		if mid == len(arr)-1 || arr[mid] > arr[mid+1] {
			k = mid
			right = mid - 1
		} else {
			left = mid + 1
		}
	}

	return k
}
