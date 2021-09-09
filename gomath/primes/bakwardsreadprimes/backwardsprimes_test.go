package bakwardsreadprimes

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func dotest(start int, stop int, exp []int) {
	var ans = BackwardsPrime(start, stop)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Tests BackwardsPrime", func() {

	It("should handle basic cases", func() {
		var a = []int{13, 17, 31, 37, 71, 73, 79, 97}
		dotest(1, 100, a)
		a = []int{13, 17, 31}
		dotest(1, 31, a)
		dotest(501, 599, nil)
	})
})
