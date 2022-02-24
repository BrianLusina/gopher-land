package factorial

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestFactorial(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Factorial Test Suite")
}

var _ = Describe("Factorial Tests", func() {
	It("should return 120 for 5!", func() {
		expected := uint64(120)
		actual := Factorial(5)
		Expect(actual).To(Equal(expected))
	})
	It("should return 3628800 for 10!", func() {
		expected := uint64(3628800)
		actual := Factorial(10)
		Expect(actual).To(Equal(expected))
	})
	It("should return 1 for 0!", func() {
		expected := uint64(1)
		actual := Factorial(0)
		Expect(actual).To(Equal(expected))
	})
	It("should return 1 for 1!", func() {
		expected := uint64(1)
		actual := Factorial(1)
		Expect(actual).To(Equal(expected))
	})
})
