package mergesort

// MergeSortConcurrent merges a list using merge sort concurrently
func MergeSortConcurrent(data []int) []int {
	if len(data) <= 1 {
		return data
	}

	doneCh := make(chan bool)
	mid := len(data) / 2
	var left []int

	go func() {
		left = MergeSortConcurrent(data[:mid])
		doneCh <- true
	}()

	right := MergeSortConcurrent(data[mid:])
	<-doneCh
	return combineLists(left, right)
}
