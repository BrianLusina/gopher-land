package lrucache

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestLruCache(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "LruCache Suite")
}

var _ = Describe("LruCache", func() {
	It("should handle put(1,1) -> put(2,2) -> get(1) -> put(3, 3) -> get(2) -> put(4, 4) -> get(1) -> get(3) -> get(4", func() {
		var lru = NewLruCache[int](2)

		lru.Put(1, 1)
		lru.Put(2, 2)

		Expect(lru.Get(1)).To(Equal(1))

		lru.Put(3, 3)

		val2 := lru.Get(2)
		Expect(val2).To(Equal(nil))

		lru.Put(4, 4)

		Expect(lru.Get(1)).To(Equal(nil))

		Expect(lru.Get(3)).To(Equal(3))
		Expect(lru.Get(4)).To(Equal(4))
	})
})
