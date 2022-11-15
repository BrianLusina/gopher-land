package cache

func main() {
	lfu := &Lfu{}
	cache := newCache(lfu)

	cache.add("a", "1")
	cache.add("b", "2")

	cache.add("c", "3")

	lru := &Lru{}
	cache.setEviction(lru)

	cache.add("d", "4")

	fifo := &Fifo{}
	cache.setEviction(fifo)

	cache.add("e", "5")
}
