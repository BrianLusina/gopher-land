package tanktruck

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
)

func dotest(h, d, vt int, exp int) {
	var ans = TankVol(h, d, vt)
	gomega.Expect(ans).To(gomega.Equal(exp))
}

func TestOrderedStream(t *testing.T) {
	var _ = ginkgo.Describe("Test Example", func() {

		ginkgo.It("should handle basic cases", func() {
			dotest(5, 7, 3848, 2940)
			dotest(2, 7, 3848, 907)
			dotest(3, 6, 3500, 1750)
			dotest(3, 6, 3501, 1750)
		})
	})
}
