package uniqueoccurences

func uniqueOccurrences(arr []int) bool {
	frequencyMap := make(map[int]int)

	for _, num := range arr {
		if _, ok := frequencyMap[num]; ok {
			frequencyMap[num]++
		} else {
			frequencyMap[num] = 1
		}
	}

	frequencySet := make(map[int]bool)
	for _, frequency := range frequencyMap {
		if _, ok := frequencySet[frequency]; !ok {
			frequencySet[frequency] = true
		}
	}

	return len(frequencySet) == len(frequencyMap)
}
