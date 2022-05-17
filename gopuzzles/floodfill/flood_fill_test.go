package floodfill

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	image    [][]int
	sr       int
	sc       int
	newColor int
	expected [][]int
}

var testCases = []testCase{
	{
		image:    [][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}},
		sr:       1,
		sc:       1,
		newColor: 2,
		expected: [][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}},
	},
	{
		image:    [][]int{{0, 0, 0}, {0, 0, 0}},
		sr:       0,
		sc:       0,
		newColor: 2,
		expected: [][]int{{2, 2, 2}, {2, 2, 2}},
	},
	{
		image:    [][]int{{0, 0, 0}, {0, 1, 1}},
		sr:       1,
		sc:       1,
		newColor: 1,
		expected: [][]int{{0, 0, 0}, {0, 1, 1}},
	},
}

func TestFloodFill(t *testing.T) {
	for _, test := range testCases {
		t.Run(fmt.Sprintf("%v", test.image), func(t *testing.T) {
			actual := floodFill(test.image, test.sr, test.sc, test.newColor)
			if !assert.Equal(t, test.expected, actual) {
				t.Errorf("got %v, want %v", actual, test.expected)
			}
		})
	}
}
