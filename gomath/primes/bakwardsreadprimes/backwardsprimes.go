package bakwardsreadprimes

import (
	"gopherland/math/primes"
	"gopherland/math/primes/sieveoferastothenes"
	"gopherland/math/utils"
)

func BackwardsPrime(start int, stop int) (result []int) {

	sievedPrimes := sieveoferastothenes.Sieve(start, stop)

	for _, prime := range sievedPrimes {
		reversedPrime := utils.ReverseInt(prime)

		if primes.IsPrime(reversedPrime) && reversedPrime != prime {
			result = append(result, prime)
		}
	}
	return
}
