package utils

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestStringUtils(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "StringUtils Test Suite")
}

var _ = Describe("StringUtils Tests", func() {

	Describe("Reverse String Test Cases", func() {
		type testCase struct {
			input    string
			expected string
		}

		testCases := []testCase{
			{"HnDMao", "oaMDnH"},
			{"ClNNxX", "XxNNlC"},
			{"iRvxxH", "HxxvRi"},
			{"bqTVvA", "AvVTqb"},
			{"wvSyRu", "uRySvw"},
		}

		for _, testCase := range testCases {
			It(fmt.Sprintf("should reverse string %s to %s", testCase.input, testCase.expected), func() {
				input := testCase.input
				expected := testCase.expected
				actual := ReverseString(input)
				Expect(actual).To(Equal(expected))
			})
		}
	})
})
