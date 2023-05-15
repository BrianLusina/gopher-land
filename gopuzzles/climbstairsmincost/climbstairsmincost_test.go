package climbstairsmincost

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestClimbStairsMinCost(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Climb Stairs Min Cost Test Suite")
}

var _ = Describe("ClimbStairsMinCost", func() {
	It("should return 15 for input of [10,15,20]", func() {
		cost := []int{10, 15, 20}
		expected := 15
		actual := minCostClimbingStairs(cost)
		Expect(actual).Should(Equal(expected))
	})

	It("should return 6 for input of [1,100,1,1,1,100,1,1,100,1]", func() {
		cost := []int{1, 100, 1, 1, 1, 100, 1, 1, 100, 1}
		expected := 6
		actual := minCostClimbingStairs(cost)
		Expect(actual).Should(Equal(expected))
	})
})
