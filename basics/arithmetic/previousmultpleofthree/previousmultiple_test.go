package previousmultpleofthree

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestPrevMultOfThree(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prev Mult Of Three Test Suite")
}

var _ = Describe("Simple tests", func() {
	It("Expected 1 to equal nil", func() {
		Expect(PrevMultOfThree(1)).To(BeNil())
	})
	It("Expected 25 to equal nil", func() {
		Expect(PrevMultOfThree(25)).To(BeNil())
	})
	It("Expected 36 to equal 36", func() {
		Expect(PrevMultOfThree(36)).To(Equal(36))
	})
	It("Expected 1244 to equal 12", func() {
		Expect(PrevMultOfThree(1244)).To(Equal(12))
	})
	It("Expected 952406 to equal 9", func() {
		Expect(PrevMultOfThree(952406)).To(Equal(9))
	})
})
