package lsproduct

import (
	"fmt"
)

// LargestSeriesProduct takes in a string of digits and the span to get the largest product from a contiguous
// series
func LargestSeriesProduct(digits string, span int) (product int64, err error) {
	if span < 0 {
		return 0, fmt.Errorf("span is negative: %d", span)
	}

	if len(digits) < span {
		return 0, fmt.Errorf("len(%s) < span: %d < %d", digits, len(digits), span)
	}

	digitsArray := make([]int64, len(digits))

	for index, digit := range digits {
		if digit < '0' || digit > '9' {
			return 0, fmt.Errorf("input %q contains non-digits", digits)
		}

		digitsArray[index] = int64(digit - '0')
	}

	for index, last := 0, len(digitsArray)-span+1; index < last; index++ {
		currentProduct := int64(1)

		for _, digit := range digitsArray[index : index+span] {
			currentProduct *= digit
		}

		if currentProduct > product {
			product = currentProduct
		}
	}

	return
}
