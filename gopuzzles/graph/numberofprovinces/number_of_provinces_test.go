package numberofprovinces

import "testing"

type testCase struct {
	input    [][]int
	expected int
}

var testCases = []testCase{
	{
		input:    [][]int{},
		expected: 0,
	},
	{
		input:    [][]int{{1, 1, 0}, {1, 1, 0}, {0, 0, 1}},
		expected: 2,
	},
	{
		input:    [][]int{{1, 0, 0}, {0, 1, 0}, {0, 0, 1}},
		expected: 3,
	},
	{
		input:    [][]int{{1, 1, 1}, {1, 1, 1}, {1, 1, 1}},
		expected: 1,
	},
}

func TestNumberOfProvinces(t *testing.T) {
	for _, testCase := range testCases {
		actual := numberOfProvinces(testCase.input)

		if actual != testCase.expected {
			t.Errorf("numberOfProvinces(%v) = %d , expected %d", testCase.input, actual, testCase.expected)
		}
	}
}

func TestNumberOfProvincesDfs(t *testing.T) {
	for _, testCase := range testCases {
		actual := numberOfProvincesDfs(testCase.input)

		if actual != testCase.expected {
			t.Errorf("numberOfProvincesDfs(%v) = %d , expected %d", testCase.input, actual, testCase.expected)
		}
	}
}

func BenchmarkNumberOfProvinces(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}

	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			numberOfProvinces(testCase.input)
		}
	}
}

func BenchmarkNumberOfProvincesDfs(b *testing.B) {
	if testing.Short() {
		b.Skip("skipping benchmark")
	}

	for i := 0; i < b.N; i++ {
		for _, testCase := range testCases {
			numberOfProvincesDfs(testCase.input)
		}
	}
}
