package sieveoferastothenes

import (
	"gopherland/math/primes"
)

// Sieve returns a slice of all primes from start to limit
func Sieve(start int, limit int) (result []int) {
	var toSieve []int

	for i := start; i <= limit; i++ {
		toSieve = append(toSieve, i)
	}

	for _, num := range toSieve {
		if primes.IsPrime(num) {
			result = append(result, num)
		}
	}

	return
}
