package repeateddnasequences

func findRepeatedDnaSequences(dnaSequence string) []string {
	toIntMap := map[rune]int{'A': 0, 'C': 1, 'G': 2, 'T': 3}
	encodedSequence := []int{}
	for _, c := range dnaSequence {
		n := toIntMap[c]
		encodedSequence = append(encodedSequence, n)
	}

	dnaSequenceSubstrLength, dnaSequenceLength := 10, len(dnaSequence) // Length of DNA sequence to check

	if dnaSequenceLength <= dnaSequenceSubstrLength {
		return []string{}
	}

	baseAEncoding := 4 // Base-4 encoding
	rollingHashValue := 0
	output := map[string]bool{}
	seenHashes := map[int]bool{}
	aK := 1 // Stores a^k for hash updates

	// Compute the initial hash using base-4 multiplication
	for i := range dnaSequenceSubstrLength {
		rollingHashValue = rollingHashValue*baseAEncoding + encodedSequence[i]
		aK *= baseAEncoding // Precompute a^k for later use in rolling hash updates
	}

	// Store the initial hash
	seenHashes[rollingHashValue] = true

	// Sliding window approach to update the hash efficiently
	for start := 1; start < dnaSequenceLength-dnaSequenceSubstrLength+1; start++ {
		// Remove the leftmost character and add the new rightmost character
		rollingHashValue = rollingHashValue*baseAEncoding - encodedSequence[start-1]*aK + encodedSequence[start+dnaSequenceSubstrLength-1]

		// If this hash has been seen_hashes before, add the corresponding substring to the output
		found := seenHashes[rollingHashValue]
		if found {
			d := dnaSequence[start : start+dnaSequenceSubstrLength]
			result := output[d]
			if !result {
				output[d] = true
			}
		} else {
			seenHashes[rollingHashValue] = true
		}
	}

	// Convert set to list before returning
	result := []string{}
	for k := range output {
		result = append(result, k)
	}
	return result
}

func findRepeatedDnaSequencesNaive(dnaSequence string) []string {
	if len(dnaSequence) <= 10 {
		return []string{}
	}

	resultSet := map[string]bool{}
	seen := map[string]bool{}

	for idx := 0; idx < len(dnaSequence); idx++ {
		if idx+10 > len(dnaSequence) {
			break
		}
		subsequence := dnaSequence[idx : idx+10]
		if len(subsequence) < 10 {
			continue
		}
		if _, ok := seen[subsequence]; ok {
			if _, ok := resultSet[subsequence]; !ok {
				resultSet[subsequence] = true
			}
		} else {
			seen[subsequence] = true
		}
	}

	result := []string{}
	for k := range resultSet {
		result = append(result, k)
	}
	return result
}
