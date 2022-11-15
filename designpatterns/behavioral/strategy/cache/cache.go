package cache

type Cache struct {
	storage     map[string]string
	eviction    Eviction
	capacity    int
	maxCapacity int
}

func newCache(e Eviction) *Cache {
	storage := make(map[string]string)
	return &Cache{
		storage:     storage,
		eviction:    e,
		capacity:    0,
		maxCapacity: 2,
	}
}

func (c *Cache) setEviction(e Eviction) {
	c.eviction = e
}

func (c *Cache) add(key, value string) {
	if c.capacity == c.maxCapacity {
		c.evict()
	}
	c.capacity++
	c.storage[key] = value
}

func (c *Cache) get(key string) {
	delete(c.storage, key)
}

func (c *Cache) evict() {
	c.eviction.evict(c)
	c.capacity--
}
