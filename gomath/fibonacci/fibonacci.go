package fibonacci

import "math"

/*
 Time complexity : O(2^N) This is the slowest way to solve the Fibonacci Sequence because it takes exponential time.
 The amount of operations needed, for each level of recursion, grows exponentially as the depth approaches N.

 Space complexity : O(N). We need space proportionate to N to account for the max size of the stack, in memory.
 This stack keeps track of the function calls to fib(N). This has the potential to be bad in cases that there isn't enough
 physical memory to handle the increasingly growing stack, leading to a StackOverflowError.
*/
func FibRecursion(n int) int {
	if n < 1 {
		return n
	}

	return FibRecursion(n-1) + FibRecursion(n-2)
}

/*
Time complexity : O(N). Each number, starting at 2 up to and including N, is visited, computed and then stored for O(1) access later on.

Space complexity : O(N). The size of the data structure is proportionate to N.
*/
func FibMemo(n int) int {
	if n <= 1 {
		return n
	}
	memoize := func(n int) int {
		cache := map[int]int{0: 0, 1: 1}

		for i := 2; i <= n; i++ {
			cache[i] = cache[i-1] + cache[i-2]
		}
		return cache[n]
	}

	return memoize(n)
}

// Fib returns a function closure that can be used to compute fibonacci numbers
func Fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return b
	}
}

/*
Intuition Using the golden ratio, a.k.a Binet's forumula: φ = (1 + sqrt(5)) / 2 ≈ 1.6180339887....

Here's a link to find out more about how the Fibonacci sequence and the golden ratio work.

We can derive the most efficient solution to this problem using only constant time and constant space!

Algorithm

Use the golden ratio formula to calculate the Nth Fibonacci number.

Time complexity : O(1). Constant time complexity since we are using no loops or recursion and the time is based on the result of
performing the calculation using Binet's formula.

Space complexity : O(1). The space used is the space needed to create the variable to store the golden ratio formula.
*/
func FibGoldenRatio(n int) int {
	golenRatio := float64((1 + math.Sqrt(5)) / 2)
	return int(math.Round(math.Pow(golenRatio, float64(n)) / math.Sqrt(5)))
}
