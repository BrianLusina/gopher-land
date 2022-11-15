package cache

type Eviction interface {
	evict(c *Cache)
}
