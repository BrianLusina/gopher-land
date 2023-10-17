package utils

import (
	"fmt"
	"gopherland/pkg/types"
	"strings"
)

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

// SliceToString converts an slice of any type to it's string representation
func SliceToString[T any](slice []T, delim string) string {
	return strings.Trim(strings.Join(strings.Split(fmt.Sprint(slice), " "), delim), "[]")
}

// Max returns the maximum value of 2 comparable types
func Max[T types.Comparable](a, b T) T {
	if a < b {
		return b
	}
	return a
}

// Min returns the minimum value of 2 comparable types
func Min[T types.Comparable](a, b T) T {
	if a < b {
		return a
	}
	return b
}

// MinMax returns the minimum and maximum of 2 values
// For example MinMax(1, 2) returns (1, 2)
func MinMax[T types.Comparable](x, y T) (min T, max T) {
	if x < y {
		return x, y
	}
	return y, x
}

// ZipPair is a tuple of 2 elements from 2 different slices
type ZipPair[T, U any] struct {
	A T
	B U
}

// Zip returns a slice of pairs/tuples of a & b in the form of []Pair{{a[0], b[0]}, ..., {a[i], b[i]}} where i is the index
// This assumes that the length of the 2 slices are equal. If there are not, an error is returned
func Zip[T, U any](a []T, b []U) ([]ZipPair[T, U], error) {
	if len(a) != len(b) {
		return nil, fmt.Errorf("Zip: arguments must be of same length")
	}

	pairs := make([]ZipPair[T, U], len(a), len(b))

	for index := range a {
		pairs[index] = ZipPair[T, U]{a[index], b[index]}
	}

	return pairs, nil
}

// ZipDiff zips slices of different lengths by leaving the Pair field to it's Zero value
func ZipDiff[T, U any](a []T, b []U) []ZipPair[T, U] {
	// identify the minimum and the maximum lengths
	lmin, lmax := MinMax(len(a), len(b))

	pairs := make([]ZipPair[T, U], lmax)

	// build tuples up to the minimum length
	for i := 0; i < lmin; i++ {
		pairs[i] = ZipPair[T, U]{a[i], b[i]}
	}

	if lmin == lmax {
		return pairs
	}

	// build tuples with one zero value for [lmin, lmax] range
	for i := lmin; i < lmax; i++ {
		p := ZipPair[T, U]{}
		if len(a) == lmax {
			p.A = a[i]
		} else {
			p.B = b[i]
		}
		pairs[i] = p
	}

	return pairs
}

// Equal compares 2 slices and checks them for equality, returning true if they are equal false otherwise
func EqualSlices[T types.Comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}

	for idx, value := range a {
		if value != b[idx] {
			return false
		}
	}
	return true
}

// All returns tree if all the elements in the slice match the given predicate, returns false otherwise
func All[T any](elements []T, predicate func(data T) bool) bool {
	for _, element := range elements {
		if !predicate(element) {
			return false
		}
	}

	return true
}

// Any returns tree if any of the elements in the slice match the given predicate, returns false otherwise
func Any[T any](elements []T, predicate func(data T) bool) bool {
	for _, element := range elements {
		if predicate(element) {
			return true
		}
	}

	return false
}
