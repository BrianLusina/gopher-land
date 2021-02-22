package frogjumps

// FrogJumps returns the minumum number of jumps a frog has to make to get to the other side
func FrogJumps(x, y, d int) int {
	jumps := 0

	// no need to jump
	if x == y {
		return jumps
	}

	distance := y - x

	if distance%d == 0 {
		jumps = distance / d
	} else {
		jumps = (distance / d) + 1
	}

	return jumps
}
