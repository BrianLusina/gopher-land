package poweroftwo

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	n        int
	expected bool
}

var testCases = []testCase{
	{n: 1, expected: true},
	{n: 16, expected: true},
	{n: 4, expected: true},
	{n: 3, expected: false},
	{n: -512, expected: false},
	{n: 123456, expected: false},
}

func TestIsPowerOfTwo(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("IsPowerOfTwo(%d)", tc.n), func(t *testing.T) {
			actual := IsPowerOfTwo(tc.n)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkIsPowerOfTwo(b *testing.B) {
	for b.Loop() {
		for _, tc := range testCases {
			b.Run(fmt.Sprintf("IsPowerOfTwo(%d)", tc.n), func(b *testing.B) {
				_ = IsPowerOfTwo(tc.n)
			})
		}
	}
}
