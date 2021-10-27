package longestsubstringwithkchars

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLongestSubKChars(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LongestSubKChars Suite")
}

var _ = Describe("LongestSubKChars", func() {
	It("should return 3 for input aaabb and k = 3", func() {
		s := "aaabb"
		k := 3
		expected := 3
		actual := longestSubstring(s, k)
		Expect(actual).To(Equal(expected))
	})

	It("should return 5 for input aababbc and k = 2", func() {
		s := "ababbc"
		k := 2
		expected := 5
		actual := longestSubstring(s, k)
		Expect(actual).To(Equal(expected))
	})
})
