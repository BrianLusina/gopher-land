package heap

// GetParentIndex returns the parent index of a node in an Array-based heap
func GetParentIndex(i int) int {
	return (i - 1) / 2
}

// GetLeftChildIndex returns the left child of a node in an Array-Based heap
func GetLeftChildIndex(i int) int {
	return 2*i + 1
}

// GetRightChildIndex returns the right child index of a node in an Array-Based heap
func GetRightChildIndex(i int) int {
	return 2*i + 2
}
