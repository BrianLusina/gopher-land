package sumofintegers

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestSumOfIntegersInString(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "SumOfIntegersInString Suite")
}

type Case struct {
	str    string
	result int
}

var cases = []Case{
	{"12.4", 16},
	{"h3ll0w0rld", 3},
	{"2 + 3 = ", 5},
	{"Our company made approximately 1 million in gross revenue last quarter.", 1},
	{"The Great Depression lasted from 1929 to 1939.", 3868},
	{"Dogs are our best friends.", 0},
	{"The30quick20brown10f0x1203jumps914ov3r1349the102l4zy dog", 3635},
}

var _ = Describe("Test Example", func() {
	for _, testCase := range cases {
		It(fmt.Sprintf("should correctly return %d for input of %s", testCase.result, testCase.str), func() {
			actual := SumOfIntegersInString(testCase.str)
			expected := testCase.result
			Expect(actual).To(Equal(expected), "The sum of Integers in '%s' should be equal to %d", testCase.str, testCase.result)
		})
	}
})
