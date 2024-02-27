package randomizedset

import (
	"gopherland/pkg/types"
	"math/rand"
)

type RandomizedSet[T types.Comparable] struct {
	containerMap map[T]int
	data         []T
}

// New creates a new randomized set
func New[T types.Comparable]() RandomizedSet[T] {
	return RandomizedSet[T]{
		containerMap: map[T]int{},
		data:         []T{},
	}
}

// Inserts an item into the set if not present. Returns true if the item was not present, false otherwise
func (rs *RandomizedSet[T]) Insert(item T) bool {
	if _, ok := rs.containerMap[item]; ok {
		return false
	}

	// add the element to the dictionary. Setting the value as the
	// length of the list will accurately point to the index of the
	// new element. (len(some_list) is equal to the index of the last item +1)
	rs.containerMap[item] = len(rs.data)
	rs.data = append(rs.data, item)
	return true
}

// Remove removes an item val from the set if present. Returns true if the item was present, false otherwise
func (rs *RandomizedSet[T]) Remove(item T) bool {
	if _, ok := rs.containerMap[item]; !ok {
		return false
	}

	// essentially, we're going to move the last element in the list into the location of the element we want to
	// remove. This is a significantly more efficient operation than the obvious solution of removing the item and
	// shifting the values of every item in the dictionary to match their new position in the list
	index := rs.containerMap[item]
	lastElement := rs.data[len(rs.data)-1]

	rs.containerMap[lastElement] = index

	rs.data[index] = lastElement
	rs.data = rs.data[:len(rs.data)-1]

	delete(rs.containerMap, item)

	return true
}

func (rs *RandomizedSet[T]) GetRandom() T {
	if len(rs.containerMap) == 0 {
		panic("randomized set is empty")
	}

	randomIndex := rand.Intn(len(rs.data))
	return rs.data[randomIndex]
}
