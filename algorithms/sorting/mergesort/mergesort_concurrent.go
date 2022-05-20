package mergesort

import "gopherland/pkg/types"

// MergeSortConcurrent merges a list using merge sort concurrently
func MergeSortConcurrent[T types.Number](data []T) []T {
	if len(data) <= 1 {
		return data
	}

	doneCh := make(chan bool)
	mid := len(data) / 2
	var left []T

	go func() {
		left = MergeSortConcurrent(data[:mid])
		doneCh <- true
	}()

	right := MergeSortConcurrent(data[mid:])
	<-doneCh
	return combineLists(left, right)
}
