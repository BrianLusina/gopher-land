package maxprofit

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMaxProfit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Max Profit Test Suite")
}

var _ = Describe("MaxProfit Tests", func() {
	It("Expected max profit from prices [7,1,5,3,6,4] to equal 5", func() {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 5
		Expect(MaxProfit(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [7,6,4,3,1] to equal 0", func() {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0
		Expect(MaxProfit(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [1,1,1,1,1] to equal 0", func() {
		prices := []int{1, 1, 1, 1, 1}
		expected := 0
		Expect(MaxProfit(prices)).To(Equal(expected))
	})
	It("Expected max profit from prices [] to equal 0", func() {
		prices := []int{}
		expected := 0
		Expect(MaxProfit(prices)).To(Equal(expected))
	})
})
