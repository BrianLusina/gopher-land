package inmem

import (
	"sync"
	"time"
)

type now func() time.Time

type Cache struct {
	mu     sync.RWMutex
	events []Event[any]
	now    now
}

type Event[T any] struct {
	Timestamp time.Time
	Data      T
}

func NewCache() *Cache {
	return &Cache{
		events: make([]Event[any], 0),
		now:    time.Now,
	}
}

func (c *Cache) AddAll(events []Event[any]) {
	c.mu.Lock()
	c.events = append(c.events, events...)
	c.mu.Unlock()
}

func (c *Cache) GetAll() []Event[any] {
	c.mu.RLock()
	defer c.mu.RUnlock()
	events := c.events
	return events
}

func (c *Cache) TimeOlderThan(now time.Time) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	for i := 0; i < len(c.events); i++ {
		if c.events[i].Timestamp.After(now) {
			// trims the events
			c.events = c.events[i:]
			return
		}
	}
}
