package sparsevector

import "testing"

type testCase struct {
	name     string
	vec1     []int
	vec2     []int
	expected int
}

var testCases = []testCase{
	{
		name:     "Test Case 1",
		vec1:     []int{1, 0, 0, 2, 3},
		vec2:     []int{0, 3, 0, 4, 0},
		expected: 8,
	},
	{
		name:     "Test Case 2",
		vec1:     []int{0, 1, 0, 0, 0},
		vec2:     []int{0, 0, 0, 0, 2},
		expected: 0,
	},
	{
		name:     "Test Case 3",
		vec1:     []int{0, 0, 0, 0, 0},
		vec2:     []int{1, 2, 3, 4, 5},
		expected: 0,
	},
	{
		name:     "Test Case 4",
		vec1:     []int{1, 2, 3, 4, 5},
		vec2:     []int{0, 0, 0, 0, 0},
		expected: 0,
	},
	{
		name:     "Test Case 5",
		vec1:     []int{1, 2, 3},
		vec2:     []int{4, 5, 6},
		expected: 32,
	},
}

func TestSparseVectorDotProduct(t *testing.T) {
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			vec1 := New(tc.vec1)
			vec2 := New(tc.vec2)
			result := vec1.DotProduct(vec2)

			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}
}

func BenchmarkSparseVectorDotProduct(b *testing.B) {
	if testing.Short() {
		b.Skip("Skipping benchmark in short mode.")
	}

	for b.Loop() {
		for _, tc := range testCases {
			vec1 := New(tc.vec1)
			vec2 := New(tc.vec2)
			vec1.DotProduct(vec2)
		}
	}
}
