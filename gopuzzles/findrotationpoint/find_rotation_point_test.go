package findrotationpoint

import "testing"

type testCase struct {
	words    []string
	expected int
}

var testCases = []testCase{
	{
		words: []string{
			"ptolemaic",
			"retrograde",
			"supplant",
			"undulate",
			"xenoepist",
			"asymptote",
			"babka",
			"banoffee",
			"engender",
			"karpatka",
			"othellolagkage",
		},
		expected: 5,
	},
	{
		words: []string{
			"cape",
			"cake",
		},
		expected: 1,
	},
	{
		words: []string{
			"grape",
			"orange",
			"plum",
			"radish",
			"apple",
		},
		expected: 4,
	},
}

func TestFindRotationPoint(t *testing.T) {
	for _, tc := range testCases {
		actual := findRotationPoint(tc.words)

		if actual != tc.expected {
			t.Fatalf("findRotationPoint(%v) = %d expected: %d", tc.words, actual, tc.expected)
		}
	}
}

func BenchmarkFindRotationPoint(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			findRotationPoint(tc.words)
		}
	}
}
