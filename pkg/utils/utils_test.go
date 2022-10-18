package utils

import (
	"fmt"
	"math/rand"
	"testing"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRangeSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Range Suite")
}

type testCase struct {
	start    int
	end      int
	step     int
	expected []int
}

var testCases = []testCase{
	{
		start:    0,
		end:      10,
		step:     1,
		expected: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
	{
		start:    1,
		end:      11,
		step:     1,
		expected: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	},
	{
		start:    0,
		end:      30,
		step:     5,
		expected: []int{0, 5, 10, 15, 20, 25},
	},
	{
		start:    1,
		end:      4,
		step:     0,
		expected: []int{1, 1, 1},
	},
	{
		start:    10,
		end:      0,
		step:     0,
		expected: nil,
	},
}

func referenceSolution(start, end, step int) []int {
	if start > end {
		return []int{}
	}
	if step == 0 {
		o := make([]int, end-start)
		for i := range o {
			o[i] = start
		}
		return o
	} else {
		o := []int{}
		for start < end {
			o = append(o, start)
			start += step
		}
		return o
	}
}

var _ = Describe("Range Test Cases", func() {

	for _, testCase := range testCases {
		start := testCase.start
		end := testCase.end
		step := testCase.step
		expected := testCase.expected

		It(fmt.Sprintf("Range(%d, %d, %d)", start, end, step), func() {
			if len(expected) == 0 {
				Expect(Range(start, end, step)).To(BeEmpty(), "With start = %d, end = %d, step = %d", start, end, step)
			} else {
				Expect(Range(start, end, step)).To(Equal(expected), "With start = %d, end = %d, step = %d", start, end, step)
			}
		})
	}

	rand.Seed(time.Now().UTC().UnixNano())

	for i := 0; i < 200; i++ {
		start := rand.Intn(100) - 50
		end := start + rand.Intn(50) - rand.Intn(4)
		step := rand.Intn(20)

		It(fmt.Sprintf("Range(%d, %d, %d)", start, end, step), func() {
			expected := referenceSolution(start, end, step)

			if len(expected) == 0 {
				Expect(Range(start, end, step)).To(BeEmpty(), "With start = %d, end = %d, step = %d", start, end, step)
			} else {
				Expect(Range(start, end, step)).To(Equal(expected), "With start = %d, end = %d, step = %d", start, end, step)
			}
		})
	}
})
