package sieveoferastothenes

import (
	"gopherland/gomath/primes"
)

func SieveOfEratosthenes(start int, limit int) (result []int) {
	toSieve := []int{}

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
