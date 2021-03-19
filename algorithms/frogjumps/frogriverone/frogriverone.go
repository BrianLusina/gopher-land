package frogriverone

// TimeFrogJump returns the earliest time that a frog can jump to the other side of the river
// Uses a Set to store unique values from leafPositions[]. When the size equals to endPosition,
// means the leaves have covered all positions from 1 to endPosition.
// So the frog can get to the position endPosition + 1.
func TimeFrogJump(endPosition int, leafPositions []int) int {
	values := make(map[int]int)

	for time := 0; time < len(leafPositions); time++ {
		values[time] = leafPositions[time]

		if len(values) == endPosition {
			return time
		}
	}

	return -1
}
