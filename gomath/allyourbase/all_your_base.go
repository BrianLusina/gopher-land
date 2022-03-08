package allyourbase

import (
	"errors"
	"math"
)

func ConvertToBase(inputBase int, inputDigits []int, outputBase int) ([]int, error) {
	if inputBase < 2 {
		return nil, errors.New("input base must be >= 2")
	}
	if outputBase < 2 {
		return nil, errors.New("output base must be >= 2")
	}

	for _, digit := range inputDigits {
		if digit < 0 || digit >= inputBase {
			return nil, errors.New("all digits must satisfy 0 <= d < input base")
		}
	}
	total := toBase(inputBase, inputDigits)
	return toDigits(total, outputBase), nil
}

func toBase(base int, digits []int) int {
	total := 0
	l := len(digits)

	for x := l - 1; x >= 0; x-- {
		total += digits[l-1-x] * int(math.Pow(float64(base), float64(x)))
	}
	return total
}

func reverse(a []int) []int {
	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}
	return a
}

func toDigits(total int, base int) []int {
	digits := []int{}
	for total >= base {
		digits = append(digits, total%base)
		total /= base
	}
	digits = append(digits, total%base)
	return reverse(digits)
}
