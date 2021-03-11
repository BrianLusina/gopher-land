package quicksort

// partitions a slice into 2 and returns the pivot index. Assumes the pivot is at the end of the slice
func partition(theSlice []int, startIndex, endIndex int) int {

	pivot := theSlice[endIndex]

	leftIndex := startIndex
	rightIndex := endIndex - 1

	for leftIndex <= rightIndex {
		// walk until we find something on the left side that belongs on the right (less than the pivot)
		for leftIndex <= endIndex && theSlice[leftIndex] < pivot {
			leftIndex++
		}

		// walk until we find something on the right side that belongs on the left(greater than or equal to the pivot)
		for rightIndex >= startIndex && theSlice[rightIndex] >= pivot {
			rightIndex--
		}

		if leftIndex < rightIndex {
			// swap the items at the left_index and right_index, moving the element that's smaller than the pivot to the left
			// # half and the element that's larger than the pivot to the right half
			theSlice[rightIndex], theSlice[leftIndex] = theSlice[leftIndex], theSlice[rightIndex]
		} else {
			// unless we have looked at all the elements in the list and are done partitioning. In that case, move the pivot element
			// into it's final position
			theSlice[endIndex], theSlice[leftIndex] = theSlice[leftIndex], theSlice[endIndex]
		}
	}

	return leftIndex
}

// quicksortSubSlice uses recurstion to sort each partition of the slice
func quicksortSubSlice(theSlice []int, startIndex, endIndex int) {

	// base case, list with 0 or 1 element
	if startIndex >= endIndex {
		return
	}

	// divide the list into 2 smaller sublists
	pivotIndex := partition(theSlice, startIndex, endIndex)

	// Recursively sort each sublist
	quicksortSubSlice(theSlice, startIndex, pivotIndex-1)
	quicksortSubSlice(theSlice, pivotIndex+1, endIndex)
}

// Quicksort sorts a slice of integers using quicksort algorithm
func Quicksort(theSlice []int) []int {
	length := len(theSlice)

	// Nothing to sort here
	if length <= 1 {
		return theSlice
	}

	quicksortSubSlice(theSlice, 0, length-1)
	return theSlice
}
