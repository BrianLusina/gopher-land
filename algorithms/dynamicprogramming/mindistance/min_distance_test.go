package mindistance

import "testing"

type testCase struct {
	a        string
	b        string
	expected int
}

var testCases = []testCase{
	{
		a:        "abcd",
		b:        "abcd",
		expected: 0,
	},
	{
		a:        "",
		b:        "abcd",
		expected: 4,
	},
	{
		a:        "abad",
		b:        "abac",
		expected: 1,
	},
	{
		a:        "Anshuman",
		b:        "Antihuman",
		expected: 2,
	},
	{
		a:        "horse",
		b:        "ros",
		expected: 3,
	},
	{
		a:        "intention",
		b:        "execution",
		expected: 5,
	},
}

func TestMinDistance(t *testing.T) {
	for _, tc := range testCases {
		actual := minDistance(tc.a, tc.b)
		if actual != tc.expected {
			t.Fatalf("minDistance(a=%s, b=%s)=%d, expected=%d", tc.a, tc.b, actual, tc.expected)
		}
	}
}

func BenchmarkMinDistance(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := minDistance(tc.a, tc.b)
			if actual != tc.expected {
				b.Fatalf("minDistance(a=%s, b=%s)=%d, expected=%d", tc.a, tc.b, actual, tc.expected)
			}
		}
	}
}
