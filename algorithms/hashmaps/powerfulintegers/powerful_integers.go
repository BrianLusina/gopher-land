package powerfulintegers

func powerfulIntegers(x int, y int, bound int) []int {
	powerfulInts := make(map[int]bool)

	for i := 1; i < bound; i *= x {
		for j := 1; i+j <= bound; j *= y {
			powerfulInts[i+j] = true
		}
	}

	result := make([]int, 0, len(powerfulInts))
	for k := range powerfulInts {
		result = append(result, k)
	}

	return result
}
