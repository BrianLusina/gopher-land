package primegap

import (
	"gopherland/math/primes"
	"gopherland/math/primes/sieveoferastothenes"
)

// Gap returns the first pair of prime numbers in the range [m, n] that are a g distance apart.
func Gap(g, m, n int) []int {
	if g < 2 || m > n || (n-m) < g {
		return nil
	}

	if (n-m) == g && (primes.IsPrime(m) && primes.IsPrime(n)) {
		return []int{m, n}
	}

	primeNumbers := sieveoferastothenes.Sieve(m, n)

	for idx := 0; idx < len(primeNumbers); idx++ {
		prime := primeNumbers[idx]
		next := prime + g

		if primes.IsPrime(next) && primeNumbers[idx+1] == next {
			return []int{prime, next}
		}
	}

	return nil
}
