package lrucache

type (
	node[T comparable] struct {
		key        any
		data       T
		prev, next *node[T]
	}

	// LruCache represents an LRU Cache
	LRUCache[T comparable] struct {
		capacity   int
		head, tail *node[T]
		lookup     map[any]*node[T]
		size       int
	}
)

// NewLruCache creates a new LRU Cache
func NewLruCache[T comparable](capacity int) *LRUCache[T] {
	m := map[any]*node[T]{}
	return &LRUCache[T]{capacity: capacity, lookup: m}
}

// Put adds a new key value pair to the cache
func (l *LRUCache[T]) Put(key any, value T) {
	if node, ok := l.lookup[key]; ok {
		node.data = value
		l.moveToLast(node)
		return
	}

	if l.size < l.capacity {
		l.append(key, value)
		l.size++
		return
	}

	node := l.head
	l.moveToLast(node)
	delete(l.lookup, node.key)

	l.lookup[key] = node
	node.key = key
	node.data = value
}

// Get returns the value for the given key, returns nil if none exists
func (l *LRUCache[T]) Get(key any) any {
	node, ok := l.lookup[key]
	if ok {
		data := node.data
		l.moveToLast(node)
		return data
	}
	return nil
}

func (l *LRUCache[T]) moveToLast(node *node[T]) {
	if node == l.tail {
		return
	}
	if node == l.head {
		l.head = l.head.next
		l.head.prev = nil
	} else {
		node.prev.next = node.next
		node.next.prev = node.prev
	}
	l.tail.next = node
	node.prev = l.tail
	l.tail = l.tail.next
	l.tail.next = nil
}

func (l *LRUCache[T]) append(key any, value T) {
	node := &node[T]{
		key:  key,
		data: value,
	}

	if l.tail == nil {
		l.tail = node
		l.head = node
	} else {
		l.tail.next = node
		node.prev = l.tail
		l.tail = node
	}
	l.lookup[key] = node
}
