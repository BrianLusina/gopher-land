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

var _ = Describe("MaxProfit 2 Pointers Tests", func() {
	It("Expected max profit from prices [7,1,5,3,6,4] to equal 5", func() {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 5
		Expect(MaxProfitTwoPointers(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [7,6,4,3,1] to equal 0", func() {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0
		Expect(MaxProfitTwoPointers(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [1,1,1,1,1] to equal 0", func() {
		prices := []int{1, 1, 1, 1, 1}
		expected := 0
		Expect(MaxProfitTwoPointers(prices)).To(Equal(expected))
	})
	It("Expected max profit from prices [] to equal 0", func() {
		prices := []int{}
		expected := 0
		Expect(MaxProfitTwoPointers(prices)).To(Equal(expected))
	})
})

var _ = Describe("MaxProfit 2 Tests", func() {
	It("Expected max profit from prices [7,1,5,3,6,4] to equal 7", func() {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 7
		Expect(MaxProfit2(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [7,6,4,3,1] to equal 0", func() {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0
		Expect(MaxProfit2(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [1,1,1,1,1] to equal 0", func() {
		prices := []int{1, 1, 1, 1, 1}
		expected := 0
		Expect(MaxProfit2(prices)).To(Equal(expected))
	})
	It("Expected max profit from prices [] to equal 0", func() {
		prices := []int{}
		expected := 0
		Expect(MaxProfit2(prices)).To(Equal(expected))
	})
})

var _ = Describe("MaxProfit 3 Tests", func() {
	It("Expected max profit from prices [7,1,5,3,6,4] to equal 7", func() {
		prices := []int{7, 1, 5, 3, 6, 4}
		expected := 7
		Expect(MaxProfit3(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [7,6,4,3,1] to equal 0", func() {
		prices := []int{7, 6, 4, 3, 1}
		expected := 0
		Expect(MaxProfit3(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [1, 2, 3, 4, 5] to equal 4", func() {
		prices := []int{1, 2, 3, 4, 5}
		expected := 4
		Expect(MaxProfit3(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [3, 3, 5, 0, 0, 3, 1, 4] to equal 6", func() {
		prices := []int{3, 3, 5, 0, 0, 3, 1, 4}
		expected := 6
		Expect(MaxProfit3(prices)).To(Equal(expected))
	})

	It("Expected max profit from prices [1,1,1,1,1] to equal 0", func() {
		prices := []int{1, 1, 1, 1, 1}
		expected := 0
		Expect(MaxProfit3(prices)).To(Equal(expected))
	})
	It("Expected max profit from prices [] to equal 0", func() {
		prices := []int{}
		expected := 0
		Expect(MaxProfit3(prices)).To(Equal(expected))
	})
})
