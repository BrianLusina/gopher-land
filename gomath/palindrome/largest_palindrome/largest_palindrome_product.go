package palindromeproduct

import (
	"gopherland/gostrings/utils"
	"strconv"
)

// largestPalindromProduct returns the largest palindrome product of two n-digit numbers.
func largestPalindromeProduct(n int) int {
	upperLimit := 0

	for i := 0; i < n; i++ {
		upperLimit *= 10
		upperLimit += 9
	}

	lowerLimit := (1 + upperLimit) / 10
	maxProduct := 0

	for x := upperLimit; x >= lowerLimit; x-- {
		for y := x; y >= lowerLimit; y-- {
			product := x * y

			if product < maxProduct {
				break
			}

			productStr := strconv.Itoa(product)

			if productStr == utils.ReverseString(productStr) && product > maxProduct {
				maxProduct = product
			}
		}
	}

	return maxProduct
}
