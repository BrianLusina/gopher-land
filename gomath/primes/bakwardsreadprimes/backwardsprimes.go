package bakwardsreadprimes

import (
	"gopherland/gomath/primes"
	"gopherland/gomath/primes/sieveoferastothenes"
	"gopherland/gomath/utils"
)

func BackwardsPrime(start int, stop int) (result []int) {

	sievedPrimes := sieveoferastothenes.SieveOfEratosthenes(start, stop)

	for _, prime := range sievedPrimes {
		reversedPrime := utils.ReverseInt(prime)

		if primes.IsPrime(reversedPrime) && reversedPrime != prime {
			result = append(result, prime)
		}
	}
	return
}
