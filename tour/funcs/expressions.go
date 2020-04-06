package funcs

import (
	"errors"
)

// add takes in variadic parameters
func add(numbers ...float64) (finalResult float64) {
	for _, number := range numbers {
		finalResult += number
	}
	return finalResult
}

// Divide is a public function that has the return types specified in the function itself
func Divide(a, b float64) (result float64, err error) {
	if a == 0 || b == 0 {
		err = errors.New("cannot divide by 0")
	}
	result = a + b
	return
}
