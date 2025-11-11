package utils

// MakeArray creates an array and fills it with provided values
// if fillValue is nil, the array will be filled with nil equivalent values for provided type
// otherwise, it will be filled with the provided value
func MakeArray[T any](size int, fillValue T) []T {
	array := make([]T, size)

	if !IsZero(&fillValue) {
		for idx := range array {
			array[idx] = fillValue
		}
	}

	return array
}

func sum2(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 2 {
		total += s[i]
	}

	return total
}

func sum8(s []int64) int64 {
	var total int64
	for i := 0; i < len(s); i += 8 {
		total += s[i]
	}

	return total
}

// EqualUnorderedSlices checks if two slices are equal, without considering the order of elements.
// It returns true if the two slices are equal, false otherwise.
func EqualUnorderedSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	counts := make(map[T]int)
	for _, item := range a {
		counts[item]++
	}
	for _, item := range b {
		counts[item]--
		if counts[item] < 0 {
			return false
		}
	}
	return true
}

// EqualMatrices checks if two 2D matrices are equal, without considering the order of elements.
// It returns true if the two matrices are equal, false otherwise.
func EqualMatrices[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if len(a[i]) != len(b[i]) {
			return false
		}
		for j := range a[i] {
			if a[i][j] != b[i][j] {
				return false
			}
		}
	}
	return true
}
