package reversestring

import "testing"

type testCase struct {
	text     string
	expected string
}

var testCases = []testCase{
	{
		text:     "!evitacudE ot emocleW",
		expected: "Welcome to Educative!",
	},
}

func TestReverseStr(t *testing.T) {
	for _, tc := range testCases {
		actual := reverseString(tc.text)
		if actual != tc.expected {
			t.Fatalf("reverseString(%s) = %s, expected %s", tc.text, actual, tc.expected)
		}

	}
}

func BenchmarkReverseStr(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := reverseString(tc.text)
			if actual != tc.expected {
				b.Fatalf("reverseString(%s) = %s, expected %s", tc.text, actual, tc.expected)
			}
		}
	}
}
