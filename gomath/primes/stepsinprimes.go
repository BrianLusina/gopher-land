package primes

func Step(g, m, n int) []int {
	if (n - m) < g {
		return nil
	}

	if (n-m) == g && (IsPrime(m) && IsPrime(n)) {
		return []int{m, n}
	}

	for x := m; x <= n; x++ {
		second := x + g

		if IsPrime(x) && IsPrime(second) {
			return []int{x, second}
		}
	}

	return nil
}
