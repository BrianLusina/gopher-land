package longestcommonsubsequence

import "testing"

type testCase struct {
	text1    string
	text2    string
	expected int
}

var testCases = []testCase{
	{
		text1:    "abcde",
		text2:    "ace",
		expected: 3,
	},
	{
		text1:    "abc",
		text2:    "abc",
		expected: 3,
	},
	{
		text1:    "abc",
		text2:    "def",
		expected: 0,
	},
}

func TestLongestCommonSubsequence(t *testing.T) {
	for _, tc := range testCases {
		actual := longestCommonSubsequence(tc.text1, tc.text2)
		if actual != tc.expected {
			t.Fatalf("longestCommonSubsequence(text1=%s, text=2%s)=%d expected %d", tc.text1, tc.text2, actual, tc.expected)
		}
	}
}

func BenchmarkLongestCommonSubsequence(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := longestCommonSubsequence(tc.text1, tc.text2)
			if actual != tc.expected {
				b.Fatalf("longestCommonSubsequence(text1=%s, text=2%s)=%d expected %d", tc.text1, tc.text2, actual, tc.expected)
			}
		}
	}
}
