package timemap

import (
	"testing"

	"github.com/onsi/ginkgo"
	"github.com/onsi/gomega"
	"github.com/stretchr/testify/assert"
)

func TestTimeMap(t *testing.T) {
	gomega.RegisterFailHandler(ginkgo.Fail)
	ginkgo.RunSpecs(t, "TimeMap Test Suite")
}

var _ = ginkgo.Describe("TimeMap", func() {
	ginkgo.It("set('foo', 'bar', 1) -> get('foo', 1) -> get('foo', 3) -> get('foo', 3) -> set('foo', 'bar2', 4) -> get('foo', 4) -> get('foo', 5)", func() {
		timeMap := New()

		// store the key "foo" and value "bar" along with timestamp = 1.
		timeMap.Set("foo", "bar", 1)

		// return "bar"
		expectedBar := "bar"
		actualAtTime1 := timeMap.Get("foo", 1)
		assert.Equal(ginkgo.GinkgoT(), expectedBar, actualAtTime1)

		// return "bar", since there is no value corresponding to foo at timestamp 3 and timestamp 2, then the only value is at timestamp 1 is "bar".
		expectedBar2 := "bar"
		actualAtTime3 := timeMap.Get("foo", 3)
		assert.Equal(ginkgo.GinkgoT(), expectedBar2, actualAtTime3)

		// store the key "foo" and value "bar2" along with timestamp = 4.
		timeMap.Set("foo", "bar2", 4)

		// # return "bar2"
		expectedBar22 := "bar2"
		actualAtTime4 := timeMap.Get("foo", 4)
		assert.Equal(ginkgo.GinkgoT(), expectedBar22, actualAtTime4)

		// # return "bar2"
		expectedBar23 := "bar2"
		actualAtTime5 := timeMap.Get("foo", 5)
		assert.Equal(ginkgo.GinkgoT(), expectedBar23, actualAtTime5)
	})
})
