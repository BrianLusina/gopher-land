package maxlengthdifference

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestMaxLengthDifference(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "MaxLengthDifference Suite")
}

var _ = Describe("Tests MxDifLg", func() {

	It("should return -1 if s1 input is empty", func() {
		var s1 []string
		s2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
		ans := MxDifLg(s1, s2)
		exp := -1
		Expect(ans).To(Equal(exp))
	})

	It("should return -1 if s2 input is empty", func() {
		s1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
		var s2 []string
		ans := MxDifLg(s1, s2)
		exp := -1
		Expect(ans).To(Equal(exp))
	})

	It("should return -1 if both inputs are empty", func() {
		var s1 []string
		var s2 []string
		ans := MxDifLg(s1, s2)
		exp := -1
		Expect(ans).To(Equal(exp))
	})

	It("should handle s1 of length 10 and s2 of length 3", func() {
		s1 := []string{"hoqq", "bbllkw", "oox", "ejjuyyy", "plmiis", "xxxzgpsssa", "xxwwkktt", "znnnnfqknaz", "qqquuhii", "dvvvwz"}
		s2 := []string{"cccooommaaqqoxii", "gggqaffhhh", "tttoowwwmmww"}
		ans := MxDifLg(s1, s2)
		exp := 13
		Expect(ans).To(Equal(exp))
	})

	It("should handle s1 with more strings than s2", func() {
		s1 := []string{"ejjjjmmtthh", "zxxuueeg", "aanlljrrrxx", "dqqqaaabbb", "oocccffuucccjjjkkkjyyyeehh"}
		s2 := []string{"bbbaaayddqbbrrrv"}
		ans := MxDifLg(s1, s2)
		exp := 10
		Expect(ans).To(Equal(exp))
	})
})
