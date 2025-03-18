package longestsubstringwithkchars

import (
	"math"
	"unicode/utf8"
)

func longestSubstringUtil(s string, start, end, k int) int {
	if end < k {
		return 0
	}

	countMap := [26]int{}

	for i := start; i < end; i++ {
		ar, _ := utf8.DecodeRuneInString("a")
		r, _ := utf8.DecodeRuneInString(string(s[i]))

		countMap[r-ar] += 1
	}

	for mid := start; mid < end; mid++ {
		utf8.DecodeRuneInString(string(s[mid]))

		ar, _ := utf8.DecodeRuneInString("a")
		r, _ := utf8.DecodeRuneInString(string(s[mid]))

		if countMap[r-ar] >= k {
			continue
		}

		midNext := mid + 1

		x, _ := utf8.DecodeRuneInString(string(s[midNext]))

		for midNext < end && countMap[x-ar] < k {
			midNext += 1
		}

		leftSub := longestSubstringUtil(s, start, mid, k)
		rightSub := longestSubstringUtil(s, midNext, end, k)

		return int(math.Max(float64(leftSub), float64(rightSub)))
	}

	return end - start
}

/**
 * Divide and Conquer is one of the popular strategies that work in 2 phases.
 *
 * Divide the problem into subproblems. (Divide Phase).
 * Repeatedly solve each subproblem independently and combine the result to solve the original problem. (Conquer Phase)
 * We could apply this strategy by recursively splitting the string into substrings and combine the result to find the
 * longest substring that satisfies the given condition. The longest substring for a string starting at index start and
 *  ending at index end can be given by,
 *
 * longestSustring(start, end) = max(longestSubstring(start, mid), longestSubstring(mid+1, end))
 * Finding the split position (mid)
 *
 * The string would be split only when we find an invalid character. An invalid character is the one with a frequency
 * of less than k. As we know, the invalid character cannot be part of the result, we split the string at the index
 * where we find the invalid character, recursively check for each split, and combine the result.
 *
 * Algorithm
 *
 * Build the countMap with the frequency of each character in the string s.
 * Find the position for mid index by iterating over the string. The mid index would be the first invalid character in
 * the string.
 * Split the string into 2 substrings at the mid index and recursively find the result.
 * To make it more efficient, we ignore all the invalid characters after the mid index as well, thereby reducing the
 * number of recursive calls.
 *
 * Complexity Analysis
 *
 * Time Complexity: O(N^2), where N is the length of string ss. Though the algorithm performs better in most cases,
 * the worst case time complexity is still (N ^ 2).
 *
 * In cases where we perform split at every index, the maximum depth of recursive call could be O(N). For each
 * recursive call it takes O(N) time to build the countMap resulting in O(n ^ 2) time complexity.
 *
 * Space Complexity: O(N) This is the space used to store the recursive call stack. The maximum depth of recursive
 * call stack would be O(N).
 *
 * @param s: String to evaluate for
 * @param k: length of the longest substring
 * @return: length of longest substring with at most repeating characters of length k
 */
func longestSubstring(s string, k int) int {
	return longestSubstringUtil(s, 0, len(s), k)
}
