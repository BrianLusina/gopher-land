package findrotationpoint

func findRotationPoint(words []string) int {
	firstWord := words[0]
	floorIndex, ceilingIndex := 0, len(words)-1

	for floorIndex < ceilingIndex {
		guessIndex := floorIndex + ((ceilingIndex - floorIndex) / 2)

		guessWord := words[guessIndex]

		if guessWord >= firstWord {
			floorIndex = guessIndex
		} else {
			ceilingIndex = guessIndex
		}

		if floorIndex+1 == ceilingIndex {
			return ceilingIndex
		}
	}

	return -1
}
