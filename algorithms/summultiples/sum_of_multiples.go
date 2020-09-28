package summultiples

// SumMultiples returns the sum of all divisors below limit excluding the limit
func SumMultiples(limit int, divisors ...int) (sum int) {
	for x := 1; x < limit; x++ {
		for _, divisor := range divisors {
			if divisor != 0 && x%divisor == 0 {
				sum += x
				break
			}
		}
	}
	return
}
