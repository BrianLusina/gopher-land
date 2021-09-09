package utils

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestReverseInt(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "ReverseInt Suite")
}

var _ = Describe("ReverseInt", func() {
	It("should return 97 for n = 79", func() {
		n := 79
		expected := 97

		actual := ReverseInt(n)
		Expect(actual).To(Equal(expected))
	})

	It("should return 9 for n = 9", func() {
		n := 9
		expected := 9

		actual := ReverseInt(n)
		Expect(actual).To(Equal(expected))
	})

	It("should return 861 for n = 168", func() {
		n := 168
		expected := 861

		actual := ReverseInt(n)
		Expect(actual).To(Equal(expected))
	})

	It("should return 675 for n = 576", func() {
		n := 675
		expected := 576

		actual := ReverseInt(n)
		Expect(actual).To(Equal(expected))
	})

	It("should return 54321 for n = 12345", func() {
		n := 12345
		expected := 54321

		actual := ReverseInt(n)
		Expect(actual).To(Equal(expected))
	})
})
