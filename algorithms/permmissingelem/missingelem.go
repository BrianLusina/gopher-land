package permmissingelem

// PermMissingElem finds the missing number from a list of numbers
func PermMissingElem(numbers []int) int {
	missingElement := len(numbers) + 1

	for index, number := range numbers {
		missingElement = missingElement ^ number ^ (index + 1)
	}
	return missingElement
}
