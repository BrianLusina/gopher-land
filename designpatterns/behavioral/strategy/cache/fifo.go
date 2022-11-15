package cache

import "fmt"

type Fifo struct {
}

func (f *Fifo) evict(c *Cache) {
	fmt.Println("Evicting using FIFO strategy")
}
