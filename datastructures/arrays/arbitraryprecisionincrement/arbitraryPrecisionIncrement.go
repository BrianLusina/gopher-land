package arbitraryprecisionincrement

func arbitraryPrecisionIncrement(input []int) []int {
	result := make([]int, len(input))
	copy(result, input)

	result[len(result)-1] += 1

	for i := len(result) - 1; i > 0; i-- {
		if result[i] != 10 {
			break
		}

		result[i] = 0
		result[i-1] += 1
	}

	if result[0] == 10 {
		result[0] = 1
		result = append(result, 0)
	}

	return result
}
