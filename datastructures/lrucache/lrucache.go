package lrucache

import (
	doublylinkedlist "gopherland/datastructures/linkedlist/doublylinked"
)

type LRUCache struct {
	Capacity   int
	LinkedList doublylinkedlist.DoublyLinkedList
}

func NewLruCache(capacity int) LRUCache {
	return LRUCache{}
}

func (this *LRUCache) Get(key int) int {
	return 0
}

func (this *LRUCache) Put(key int, value int) {

}
