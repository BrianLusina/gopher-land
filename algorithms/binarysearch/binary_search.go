package binarysearch

func SearchInts(list []int, key int) int {
	if len(list) == 0 {
		return -1
	}

	// here we assume the list is sorted,
	// however, in the event it's not sorted, we need to run this sorter
	// the challenge here is that this then becomes an O(n) operation
	// which slows down the algorithm for large lists
	//sort.Ints(list)

	low, high := 0, len(list)-1

	for low <= high {
		mid := (low + high) / 2
		if list[mid] == key {
			return mid
		} else if list[mid] < key {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return -1
}
