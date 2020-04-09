package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var cache = map[int]Book{}

// random generator
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	waitGroup := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}

	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1

		// number of tasks to wait on
		waitGroup.Add(2)

		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryCache(id, m); ok {
				fmt.Println("From Cache")
				fmt.Println(b)
			}

			waitGroup.Done()
		}(id, waitGroup, mutex)

		go func(id int, waitGroup *sync.WaitGroup, m *sync.RWMutex) {
			if b, ok := queryDatabase(id, m); ok {
				fmt.Println("From DB")
				fmt.Println(b)
			}
			waitGroup.Done()
		}(id, waitGroup, mutex)
	}

	waitGroup.Wait()
}

func queryCache(id int, m *sync.RWMutex) (Book, bool) {
	// lock resource so that this goroutine has access to resource
	m.RLock()
	b, ok := cache[id]

	// unlock resource
	m.RUnlock()
	return b, ok
}

func queryDatabase(id int, m *sync.RWMutex) (Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range books {
		if b.ID == id {
			m.Lock()
			cache[id] = b
			m.Unlock()
			return b, true
		}
	}

	return Book{}, false
}
