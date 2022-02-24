package rotateformax

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRotateForMax(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RotateForMax Test Suite")
}

func dotest(n int64, exp int64) {
	var ans = MaxRot(n)
	Expect(ans).To(Equal(exp))
}

var _ = Describe("Basic Tests", func() {

	It("should handle basic cases", func() {
		dotest(38458215, 85821534)
		dotest(195881031, 988103115)
		dotest(896219342, 962193428)
	})
})
