package singlenumber

// singleNumber Finds the single number that has not been duplicated.
// If we take XOR of a number and a zero, the result will be that number, i.e. a⊕0=a. If we take XOR of two same numbers, it will return 0, i.e. a⊕a=0. If we take XOR of multiple numbers, the order doesn't affect the result, i.e. a⊕b⊕c=a⊕c⊕b.
// Therefore, if we take XOR of all numbers, what's left would be that single number as every element that appears twice would be cancelled out. For example, nums=4,1,2,1,2, we can reorder it like 1,1,2,2,4, and we got (1⊕1)⊕(2⊕2)⊕4=4.
func singleNumber(nums []int) int {
	x := 0

	for _, num := range nums {
		x ^= num
	}

	return x
}

// singleNumberMath Finds the single number that has not been duplicated. 2∗sumOfSet−(SumOfNumbers)=answer
// For example, nums=4,1,2,1,2, sumOfSet is 1+2+4=7 and sumOfNumbers is 1+1+2+2+4=10. Then the answer is 2∗7−10=4.
func singleNumberMath(nums []int) int {
	hashSet := map[int]bool{}
	hashSetSum := 0
	numSum := 0

	for _, num := range nums {
		if _, ok := hashSet[num]; !ok {
			hashSet[num] = true
			hashSetSum += num
		}
	}

	for _, num := range nums {
		numSum += num
	}

	return 2*hashSetSum - numSum
}
