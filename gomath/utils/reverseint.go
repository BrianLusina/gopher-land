package utils

// ReverseInt reverses an integer
func ReverseInt(n int) int {
	var reversed int
	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}
	return reversed
}
