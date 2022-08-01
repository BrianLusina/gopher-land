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
