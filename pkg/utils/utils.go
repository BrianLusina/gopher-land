package utils

// GetZero Value returns the zero value for a generic typ
func GetZeroValue[T any]() T {
	var result T
	return result
}

// Zero uses new built in function to return zero value
func Zero[T any]() T {
	return *new(T)
}

// IsZero is used to check if a value is a zero value(aka nil) with generics. Useful for comparables
func IsZero[T comparable](v T) bool {
	return v == *new(T)
}

// Range returns a slice of integers starting at start to end(exclusive), incremented by step
func Range(start, end, step int) []int {
	list := []int{}
	for x := start; x < end && start < end; start += step {
		list = append(list, start)
		x++
	}
	return list
}

// Sum returns the sum of a slice of integers
func Sum(numbers []int) int {
	result := 0
	for _, number := range numbers {
		result += number
	}
	return result
}
