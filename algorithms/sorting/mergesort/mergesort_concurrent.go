package mergesort

import (
	"gopherland/pkg/types"
	"sync"
)

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

func MergeSortParallelV1[T types.Comparable](data []T) {
	if len(data) <= 1 {
		return
	}

	middle := len(data) / 2
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		MergeSortParallelV1(data[:middle])
	}()

	go func() {
		defer wg.Done()
		MergeSortParallelV1(data[middle:])
	}()

	wg.Wait()
	// MergeSort()
}

const max = 2048

// MergeSortParallelV2 uses the merge sort algorithm to sort a collection of comparable data.
// This uses a threshold to determine whether to sort the list using parallelism with goroutines or
// sequentially. The threshold is set to 2048, if the data size is less than this, then Sequential
// merge sort is used, otherwise, parallel merge sort is used.
// Benchmarking this on 10,000 elements results in it being faster than V1 and sequential execution
func MergeSortParallelV2[T types.Comparable](data []T) {
	if len(data) <= 1 {
		return
	}

	if len(data) <= max {
		MergeSort(data)
	} else {
		middle := len(data) / 2

		var wg sync.WaitGroup
		wg.Add(2)

		go func() {
			defer wg.Done()
			MergeSortParallelV2(data[:middle])
		}()

		go func() {
			defer wg.Done()
			MergeSortParallelV2(data[middle:])
		}()

		wg.Wait()
		// merge
	}
}
