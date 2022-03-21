package prime

func Factors(n int64) []int64 {
	c, factors := int64(2), []int64{}
	for c*c <= n {
		for n%c == 0 {
			factors = append(factors, c)
			n /= c
		}
		c++
	}

	if n > 1 {
		factors = append(factors, n)
	}

	return factors
}
