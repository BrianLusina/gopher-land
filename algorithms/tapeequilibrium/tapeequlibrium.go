package tapeequilibrium

import "math"

// TapeEquilibrium finds the minimum difference from 2 parts of a tape
func TapeEquilibrium(tape []int) (difference int) {
	parts := []int{}

	parts[0] = tape[0]

	for index := range tape {
		parts[index+1] = tape[index+1] + parts[index]
	}

	// Arbitrarily large number
	difference = 1844674407370955161

	for i := 0; i < len(parts)-1; i++ {
		result := parts[len(parts)-1] - 2*parts[i]

		difference = int(math.Min(float64(difference), math.Abs(float64(result))))
	}

	return
}
