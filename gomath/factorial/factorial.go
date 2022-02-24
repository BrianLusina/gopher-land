package factorial

// Factorial computes the factorial of a number
func Factorial(n uint64) uint64 {
	if n <= 1 {
		return 1
	}
	return n * Factorial(n-1)
}
