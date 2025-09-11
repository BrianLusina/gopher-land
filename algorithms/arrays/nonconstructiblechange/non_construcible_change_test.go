package nonconstructiblechange

import "testing"

type testCase struct {
	name   string
	input  []int
	output int
}

var testCases = []testCase{
	{
		name:   "should return 20 for input of [5, 7, 1, 1, 2, 3, 22]",
		input:  []int{5, 7, 1, 1, 2, 3, 22},
		output: 20,
	},
	{
		name:   "should return 6 for input of [1, 1, 1, 1, 1]",
		input:  []int{1, 1, 1, 1, 1},
		output: 6,
	},
	{
		name:   "should return 55 for input of [1, 5, 1, 1, 1, 10, 15, 20, 100]",
		input:  []int{1, 5, 1, 1, 1, 10, 15, 20, 100},
		output: 55,
	},
	{
		name:   "should return 3 for input of [6, 4, 5, 1, 1, 8, 9]",
		input:  []int{6, 4, 5, 1, 1, 8, 9},
		output: 3,
	},
	{
		name:   "should return 1 for input of []",
		input:  []int{},
		output: 1,
	},
	{
		name:   "should return 1 for input of [87]",
		input:  []int{87},
		output: 1,
	},
	{
		name:   "should return 32 for input of [5, 6, 1, 1, 2, 3, 4, 9]",
		input:  []int{5, 6, 1, 1, 2, 3, 4, 9},
		output: 32,
	},
	{
		name:   "should return 19 for input of [5, 6, 1, 1, 2, 3, 43]",
		input:  []int{5, 6, 1, 1, 2, 3, 43},
		output: 19,
	},
	{
		name:   "should return 3 for input of [1, 1]",
		input:  []int{1, 1},
		output: 3,
	},
	{
		name:   "should return 1 for input of [2]",
		input:  []int{2},
		output: 1,
	},
	{
		name:   "should return 2 for input of [1]",
		input:  []int{1},
		output: 2,
	},
	{
		name:   "should return 87 for input of [109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4]",
		input:  []int{109, 2000, 8765, 19, 18, 17, 16, 8, 1, 1, 2, 4},
		output: 87,
	},
	{
		name:   "should return 29 for input of [1, 2, 3, 4, 5, 6, 7]",
		input:  []int{1, 2, 3, 4, 5, 6, 7},
		output: 29,
	},
}

func TestNonConstructibleChange(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			actual := nonConstructibleChange(tt.input)
			if actual != tt.output {
				t.Fatalf("nonConstructibleChange(%v) = %d, expected %d", tt.input, actual, tt.output)
			}
		})
	}
}

func TestNonConstructibleChangeV2(t *testing.T) {
	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			actual := nonConstructibleChangeV2(tt.input)
			if actual != tt.output {
				t.Fatalf("nonConstructibleChangeV2(%v) = %d, expected %d", tt.input, actual, tt.output)
			}
		})
	}
}

func BenchmarkNonConstructibleChange(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			nonConstructibleChange(tc.input)
		}
	}
}

func BenchmarkNonConstructibleChangeV2(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for b.Loop() {
		for _, tc := range testCases {
			nonConstructibleChangeV2(tc.input)
		}
	}
}
