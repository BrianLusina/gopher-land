package subsequence

// IsValidSubsequence checks if the sequence is a valid subsequence of the array
func IsValidSubsequence(array []int, sequence []int) bool {
	if len(array) == 0 || len(sequence) > len(array) {
		return false
	}
	i, j := 0, 0

	for i < len(sequence) && j < len(array) {
		if sequence[i] == array[j] {
			i += 1
		}
		j += 1
	}
	return i == len(sequence)
}

func IsValidSubsequenceV2(array []int, sequence []int) bool {
	if len(sequence) > len(array) || len(array) == 0 || len(sequence) == 0 {
		return false
	}

	seek := 0

	for i := range array {
		if sequence[seek] == array[i] {
			seek += 1
		}
		if seek == len(sequence) {
			return true
		}
	}

	return false
}
