package longestrepeatingcharreplacement

// characterReplacement finds the lenght of the longest substring with repeating characters after k replacements.
// This uses a sliding window technique to determine the length of the longest substring after performing at most k
// replacements.

// We use a frequency map to store a map of characters to their frequencies in the window(note that this is not in the
// whole string) but only within the sliding window.

// The objective is to find the longest valid window. So, whenever a valid window is found, we expand its size by
// moving the right pointer forward. As we move the pointer forward, we update the frequency_map as well. The frequency
// map helps us keep track of the character that appears most frequently in the window. We compare the frequency of the
// newly added character with the maximum frequency of any character seen so far - max_frequency. We update
// max_frequency when we find a new maximum.

// The window size increases only when max_frequency finds a new maximum. For this, we always want the following
// condition to hold true:

// window_size − max_frequency <= k

// We stop moving the right pointer forward, or in other words, stop expanding the window when it becomes invalid. Say
// the size of the window when it becomes invalid is l. We know the previous window with the size l−1 was valid. So, we
// move the prior window of length l−1 toward the right. To do so, the left pointer moves one step further. Remember
// that the right pointer had already moved, so we don't need to move the right pointer again.

// At this point, the last valid window has moved one step to the right, but it might still be invalid. As explained
// earlier, we are only interested in larger windows, so we don't need to decrease the window size. We move the window
// of size i−1 further and further to the right until it becomes valid again.

// If we come across a valid window, we try to expand it as much as possible, and the process continues until the right
// pointer reaches the rightmost alphabet of the string. At this point, the size of the window indicates the longest
// valid substring seen yet.

// Complexity Analysis:
// Time: O(n) where n is the length of the input string. We will have to iterate through each character in the input
// string.
// Space: O(n) as we are using a dictionary to store the character frequencies. This approach requires an auxiliary
// frequency map. The maximum number of keys in the map equals the number of unique characters in the string. If there
// are m unique characters, then the memory required is proportional to m. So the space complexity is O(m). Considering
// uppercase English letters only, m=26
func characterReplacement(s string, k int) int {
	frequencyMap := map[byte]int{}
	start := 0
	maxFrequency := 0
	result := 0

	for end := 0; end < len(s); end++ {
		char := s[end]

		if frequency, ok := frequencyMap[char]; ok {
			frequency += 1
			frequencyMap[char] = frequency
		} else {
			frequencyMap[char] = 1
		}

		maxFrequency = max(maxFrequency, frequencyMap[char])

		isValid := end+1-start-maxFrequency <= k
		if !isValid {
			frequencyMap[s[start]] -= 1
			start += 1
		}

		result = end + 1 - start
	}

	return result
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
