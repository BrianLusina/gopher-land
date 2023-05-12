package mergeintervals

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(intervals [][]int, exp [][]int) {
	var ans = MergeIntervals(intervals)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests", func() {

	It("should handle basic cases", func() {
		dotest([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}})
	})
})
