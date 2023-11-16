package maxprofit

import (
	"fmt"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testCase struct {
	prices   []int
	expected int
}

var testCasesMaxProfit1 = []testCase{
	{
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 5,
	},
	{
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	},
	{
		prices:   []int{1, 1, 1, 1, 1},
		expected: 0,
	},
	{
		prices:   []int{},
		expected: 0,
	},
}

var testCasesMaxProfit3 = []testCase{
	{
		prices:   []int{7, 1, 5, 3, 6, 4},
		expected: 7,
	},
	{
		prices:   []int{7, 6, 4, 3, 1},
		expected: 0,
	},
	{
		prices:   []int{1, 2, 3, 4, 5},
		expected: 4,
	},
	{
		prices:   []int{3, 3, 5, 0, 0, 3, 1, 4},
		expected: 6,
	},
	{
		prices:   []int{1, 1, 1, 1, 1},
		expected: 0,
	},
	{
		prices:   []int{},
		expected: 0,
	},
}

type testCaseWithFee struct {
	testCase
	fee int
}

var testCasesMaxProfitWithFee = []testCaseWithFee{
	{
		testCase: testCase{
			prices:   []int{},
			expected: 0,
		},
		fee: 0,
	},
	{
		testCase: testCase{
			prices:   []int{1, 1, 1, 1, 1},
			expected: 0,
		},
		fee: 0,
	},
	{
		testCase: testCase{
			prices:   []int{1},
			expected: 0,
		},
		fee: 1,
	},
	{
		testCase: testCase{
			prices:   []int{1, 3, 2, 8, 4, 9},
			expected: 8,
		},
		fee: 2,
	},
	{
		testCase: testCase{
			prices:   []int{1, 3, 7, 5, 10, 3},
			expected: 6,
		},
		fee: 3,
	},
	{
		testCase: testCase{
			prices:   []int{1, 4, 6, 2, 8, 3, 10, 14},
			expected: 13,
		},
		fee: 3,
	},
}

func TestMaxProfit(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Max Profit Test Suite")
}

var _ = Describe("MaxProfit Tests", func() {
	Context("MaxProfit", func() {
		for _, tc := range testCasesMaxProfit1 {
			It(fmt.Sprintf("Expected max profit from prices %v to equal %d", tc.prices, tc.expected), func() {
				Expect(MaxProfit(tc.prices)).To(Equal(tc.expected))
			})
		}
	})

	Context("MaxProfit with 2 pointers", func() {
		for _, tc := range testCasesMaxProfit1 {
			It(fmt.Sprintf("Expected max profit from prices %v to equal %d", tc.prices, tc.expected), func() {
				Expect(MaxProfitTwoPointers(tc.prices)).To(Equal(tc.expected))
			})
		}
	})

	Context("MaxProfit 2", func() {
		for _, tc := range testCasesMaxProfit1 {
			It(fmt.Sprintf("Expected max profit from prices %v to equal %d", tc.prices, tc.expected), func() {
				Expect(MaxProfit2(tc.prices)).To(Equal(tc.expected))
			})
		}
	})

	Context("MaxProfit 3", func() {
		for _, tc := range testCasesMaxProfit3 {
			It(fmt.Sprintf("Expected max profit from prices %v to equal %d", tc.prices, tc.expected), func() {
				Expect(MaxProfit3(tc.prices)).To(Equal(tc.expected))
			})
		}
	})

	Context("MaxProfit With Fee", func() {
		for _, tc := range testCasesMaxProfitWithFee {
			It(fmt.Sprintf("Expected max profit from prices %v and fee %d to equal %d", tc.prices, tc.fee, tc.expected), func() {
				Expect(MaxProfitWithFee(tc.prices, tc.fee)).To(Equal(tc.expected))
			})
		}
	})
})

func BenchmarkMaxProfitWithFee(b *testing.B) {
	if testing.Short() {
		b.Skip()
	}

	for i := 0; i < b.N; i++ {
		for _, tc := range testCasesMaxProfitWithFee {
			actual := MaxProfitWithFee(tc.prices, tc.fee)
			if actual != tc.expected {
				b.Fatalf("MaxProfitWithFee(prices=%v, fee=%d) = %d, expected %d", tc.prices, tc.fee, actual, tc.expected)
			}
		}
	}
}
