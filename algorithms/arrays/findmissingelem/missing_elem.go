package findmissingelem

// FindMissingElemXor finds the missing number from a list of numbers
func FindMissingElemXor(numbers []int) int {
	missingElement := len(numbers)

	for index, number := range numbers {
		missingElement = missingElement ^ number ^ index
	}
	return missingElement
}

func FindMissingElemSumOfNTerms(numbers []int) int {
	n := len(numbers)
	sumOfNTerms := n * (n + 1) / 2

	for _, number := range numbers {
		sumOfNTerms -= number
	}
	return sumOfNTerms
}

func FindMissingElem2(numbers []int) int {
	n := len(numbers)
	index := 0

	for index < n {
		value := numbers[index]

		if value < n && numbers[value] != value {
			numbers[index], numbers[value] = numbers[value], numbers[index]
		} else {
			index++
		}
	}

	for idx := range numbers {
		if numbers[idx] != idx {
			return idx
		}
	}
	return n
}
