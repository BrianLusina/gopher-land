package birdwatcher

func chunkSlice(slice []int, chunkSize int) [][]int {
	var chunks [][]int

	for i := 0; i < len(slice); i += chunkSize {
		end := i + chunkSize

		// necessary to avoid slicing beyond the length of the slice
		if end > len(slice) {
			end = len(slice)
		}

		chunks = append(chunks, slice[i:end])
	}

	return chunks
}

// TotalBirdCount return the total bird count by summing
// the individual day's counts.
func TotalBirdCount(birdsPerDay []int) int {
	total := 0
	for _, count := range birdsPerDay {
		total += count
	}
	return total
}

// BirdsInWeek returns the total bird count by summing
// only the items belonging to the given week.
func BirdsInWeek(birdsPerDay []int, week int) int {
	chunks := chunkSlice(birdsPerDay, 7)

	return TotalBirdCount(chunks[week-1])
	// OR
	// offset := 7 * (week - 1)
	// return TotalBirdCount(birdsPerDay[offset : offset+7])
}

// FixBirdCountLog returns the bird counts after correcting
// the bird counts for alternate days.
func FixBirdCountLog(birdsPerDay []int) []int {
	for i := 0; i < len(birdsPerDay); i += 2 {
		birdsPerDay[i] += 1
	}
	return birdsPerDay
}
