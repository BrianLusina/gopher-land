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
