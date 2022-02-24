package rotateformax

import (
	"math"
	"strconv"
)

// numberOfDigits returns the number of digits
func numberOfDigits(n int64) int {
	return len(strconv.Itoa(int(n)))
}

// MaxRot returns the maximum number from rotating the given number
func MaxRot(n int64) int64 {
	numbers := []int64{n}
	rotated := n
	max := n
	k := numberOfDigits(n)
	pow := int64(math.Pow10(k))

	for i := 0; i < k; i++ {
		firstDigit, _ := strconv.Atoi(strconv.Itoa(int(n))[0])

		left := (n*10 + firstDigit - (firstDigit * pow * 10))
		rotated = left
		numbers = append(numbers, rotated)
		max = int64(math.Max(float64(max), float64(rotated)))
	}
	return max
}
