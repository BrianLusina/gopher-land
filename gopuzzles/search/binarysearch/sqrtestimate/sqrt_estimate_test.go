package sqrtestimate

import (
	"fmt"
	"gopherland/pkg/utils"
	"math"
	"testing"
)

type testCase struct {
	n        int
	expected int
}

func generateTests() []testCase {
	start := 0
	ceil := int(math.Pow(2, 10))
	numbers := utils.Range(start, ceil, 1)

	testCases := []testCase{}

	for _, num := range numbers {
		expected := int(math.Floor(math.Sqrt(float64(num))))

		testCase := testCase{
			n:        num,
			expected: expected,
		}

		testCases = append(testCases, testCase)
	}

	return testCases
}

func TestSqrtEstimate(t *testing.T) {
	for _, tc := range generateTests() {
		t.Run(fmt.Sprintf("sqrtEstimate(%d)", tc.n), func(t *testing.T) {
			actual := sqrtEstimate(tc.n)

			if actual != tc.expected {
				t.Fatalf("expected %d, got %d", tc.expected, actual)
			}
		})
	}
}

func BenchmarkSqrtEstimate(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range generateTests() {
			b.Run(fmt.Sprintf("sqrtEstimate(%d)", tc.n), func(b *testing.B) {
				actual := sqrtEstimate(tc.n)

				if actual != tc.expected {
					b.Fatalf("expected %d, got %d", tc.expected, actual)
				}
			})
		}
	}
}
