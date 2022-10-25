package issubsequence

func IsSubsequence(s string, t string) bool {
	if len(s) > len(t) {
		return false
	}

	if len(s) == 0 {
		return true
	}

	subsequence := 0

	for i, j := 0, 0; i < len(t) && j < len(s); i++ {
		if s[j] == t[i] {
			j++
			subsequence++
		}
	}
	return subsequence == len(s)
}
