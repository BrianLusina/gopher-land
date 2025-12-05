package courseschedule

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCase struct {
	numCourses    int
	prerequisites [][]int
	expected      bool
}

var testCases = []testCase{
	{
		numCourses:    2,
		prerequisites: [][]int{{1, 0}},
		expected:      true,
	},
	{
		numCourses:    4,
		prerequisites: [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}},
		expected:      true,
	},
	{
		numCourses:    1,
		prerequisites: [][]int{},
		expected:      true,
	},
	{
		numCourses:    3,
		prerequisites: [][]int{{1, 0}, {2, 1}},
		expected:      true,
	},
	{
		numCourses:    3,
		prerequisites: [][]int{{1, 0}, {2, 1}, {1, 2}},
		expected:      false,
	},
	{
		numCourses:    3,
		prerequisites: [][]int{{1, 0}, {2, 1}, {4, 3}},
		expected:      true,
	},
}

func TestCanFinish(t *testing.T) {
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("CanFinish(%v, %v)", tc.numCourses, tc.numCourses), func(t *testing.T) {
			actual := CanFinish(tc.numCourses, tc.prerequisites)
			assert.Equal(t, tc.expected, actual)
		})
	}
}
