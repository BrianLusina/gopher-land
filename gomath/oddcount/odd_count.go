// oddcount - Count the number of odd numbers in a sequence of numbers.
package oddcount

// OddCount finds the number of odd numbers in the range [1,n).
func OddCount(n int) int {
	if n%2 != 0 {
		n -= 2
	}

	return ((n - 1) / 2) + 1
}

// OddCountV2 is the same as OddCount, but divdes by 2.
func OddCountV2(n int) int {
	return n / 2
}

// OddCountV3 is the same as OddCount, but uses bitwise operators.
func OddCountV3(n int) int {
	return n >> 1
}
