// Package smallestinfiniteset contains the SmallestInfiniteSet that is used to retrieve the smallest number in the set
// of all possible positive integers
// Complexity Analysis
// Here, n is the number of add_back(num) calls and m is the number of pop_smallest() calls
// - Time Complexity O((m+n) * log(n))
//   - In each pop_smallest() method call, in the worst case, we will need to remove a number from the hash set which
//     will take O(1) time, and the top of the min-heap which will take O(log(n)) time. Thus, for m calls it will take
//     O(m * log(n)) time.
//   - In each add_back(num) method call, we might push num in the hash set which will take O(1) time and min-heap
//     which will take O(log(n)). Thus, for n calls it will take O(n * log(n)) time.
//
// - Space complexity: O(n)
//   - In the worst case, we might add n elements in the hash set and the min-heap. Thus, it will take O(n) space
package smallestinfiniteset

import "container/heap"

// SmallestInfiniteSet contains the structure of obtaining the current smallest element in a range ordered by value
type SmallestInfiniteSet struct {
	elementSet     map[int]bool
	addedElements  *Heap
	currentElement int
}

// New creates a new SmallestInfiniteSet
func New() SmallestInfiniteSet {
	addedElements := &Heap{}
	heap.Init(addedElements)

	return SmallestInfiniteSet{
		elementSet:     map[int]bool{},
		currentElement: 1,
		addedElements:  addedElements,
	}
}

// PopSmallest pops the smallest element from the set and returns it
func (sit *SmallestInfiniteSet) PopSmallest() int {
	var result int

	// if we have elements in the heap, the top element is removed and returned
	if sit.addedElements.Len() != 0 {
		// pop the smallest element from the list
		result = heap.Pop(sit.addedElements).(int)
		// remove it from the set
		delete(sit.elementSet, result)
	} else {
		// the smallest element is the current element, we return it and set the current element to the next element
		result = sit.currentElement
		sit.currentElement++
	}

	return result
}

func (sit *SmallestInfiniteSet) AddBack(num int) {
	// if the num is greater than or equal to the current element, it is not added to the set, as it is greater than
	// our current element, this means that in the set of all positive integers, it already exists. Similarly, if it
	// is already in the element set, we do not add it as well as it might have been added before; this could be the
	// case where it is less than the current element.
	if sit.currentElement <= num || sit.elementSet[num] {
		return
	}
	heap.Push(sit.addedElements, num)
	sit.elementSet[num] = true
}
