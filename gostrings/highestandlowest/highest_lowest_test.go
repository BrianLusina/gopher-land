package highestandlowest

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestHigestLowest(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "HighestLowest Suite")
}

var _ = Describe("Example Tests", func() {
	It("should test that the solution returns the correct value", func() {
		Expect(HighAndLow("8 3 -5 42 -1 0 0 -9 4 7 4 -4")).To(Equal("42 -9"))
	})
	It("test 2", func() {
		Expect(HighAndLow("1 2 3")).To(Equal("3 1"))
	})
})
