package threepointers

import "testing"

type testCase struct {
	a, b, c  []int
	expected int
}

var testCases = []testCase{
	{
		a:        []int{1, 4, 10},
		b:        []int{2, 15, 20},
		c:        []int{10, 12},
		expected: 5,
	},
	{
		a:        []int{20, 24, 100},
		b:        []int{2, 19, 22, 79, 800},
		c:        []int{10, 12, 23, 24, 119},
		expected: 2,
	},
}

func TestMinimize(t *testing.T) {
	for _, tc := range testCases {
		actual := minimize(tc.a, tc.b, tc.c)
		if actual != tc.expected {
			t.Fatalf("minimize(a=%v, b=%v, c=%v)=%d, expected=%d", tc.a, tc.b, tc.c, actual, tc.expected)
		}
	}
}

func BenchmarkMinimize(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCases {
			actual := minimize(tc.a, tc.b, tc.c)
			if actual != tc.expected {
				b.Fatalf("minimize(a=%v, b=%v, c=%v)=%d, expected=%d", tc.a, tc.b, tc.c, actual, tc.expected)
			}
		}
	}
}
