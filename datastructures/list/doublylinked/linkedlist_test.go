package doublylinkedlist

import (
	"fmt"
	"gopherland/datastructures/list"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDeleteNodeByData(t *testing.T) {
	dll := New[int]()

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)
	dll.Append(5)

	assert.Equal(t, 5, dll.Size())

	dll.DeleteNodeByData(3)

	assert.Equal(t, 4, dll.Size())

	assert.Equal(t, "1 <-> 2 <-> 4 <-> 5 <-> nil", fmt.Sprintf("%v", dll))
}

func TestPrependOperations(t *testing.T) {
	t.Run("should successfully add a new node to a non empty list", func(t *testing.T) {
		dll := New[int]()

		dll.Append(1)
		dll.Append(2)
		dll.Append(3)
		dll.Append(4)
		dll.Append(5)

		assert.Equal(t, 5, dll.Size())

		dll.Prepend(0)

		assert.Equal(t, 6, dll.Size())

		assert.Equal(t, "0 <-> 1 <-> 2 <-> 3 <-> 4 <-> 5 <-> nil", fmt.Sprintf("%v", dll))
	})

	t.Run("should successfully add a new node to the head when the node has no head", func(t *testing.T) {
		dll := New[int]()

		assert.Equal(t, 0, dll.Size())
		dll.Prepend(1)

		assert.Equal(t, 1, dll.Size())
		assert.Equal(t, 1, dll.Head.Data)
	})
}

func TestGetMiddleNode(t *testing.T) {
	t.Run("should return nil for an empty list", func(t *testing.T) {
		dll := New[int]()
		val, err := dll.GetMiddleNode()
		assert.ErrorIs(t, err, list.ErrEmptyList)
		assert.Nil(t, val)
	})

	t.Run("should return 3 for linked list of 1 -> 2 -> 3 -> 4 -> 5", func(t *testing.T) {
		dll := New[int]()
		dll.Append(1)
		dll.Append(2)
		dll.Append(3)
		dll.Append(4)
		dll.Append(5)

		val, err := dll.GetMiddleNode()

		assert.Nil(t, err)
		assert.NotNil(t, val)
		assert.Equal(t, 3, val.Data)
	})

	t.Run("should return 7 for linked list of 2->4->6->7->5->1", func(t *testing.T) {
		dll := New[int]()
		dll.Append(2)
		dll.Append(4)
		dll.Append(6)
		dll.Append(7)
		dll.Append(5)
		dll.Append(1)

		val, err := dll.GetMiddleNode()

		assert.Nil(t, err)
		assert.NotNil(t, val)
		assert.Equal(t, 7, val.Data)
	})
}

func TestReverse(t *testing.T) {
	t.Run("should reverse a non empty linked list 1 -> 2 -> 3 -> 4 -> 5", func(t *testing.T) {
		doubly_linked_list := New[int]()
		doubly_linked_list.Append(1)
		doubly_linked_list.Append(2)
		doubly_linked_list.Append(3)
		doubly_linked_list.Append(4)
		doubly_linked_list.Append(5)

		doubly_linked_list.Reverse()

		assert.Equal(t, 5, doubly_linked_list.Head.Data)
	})

	t.Run("should reverse linked list of 2->4->6->7->5->1", func(t *testing.T) {
		doubly_linked_list := New[int]()
		doubly_linked_list.Append(2)
		doubly_linked_list.Append(4)
		doubly_linked_list.Append(6)
		doubly_linked_list.Append(7)
		doubly_linked_list.Append(5)
		doubly_linked_list.Append(1)

		doubly_linked_list.Reverse()

		assert.Equal(t, 1, doubly_linked_list.Head.Data)
	})
}

func TestRotate(t *testing.T) {

	t.Run("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 when k=0", func(t *testing.T) {
		dll := New[int]()
		dll.Append(10)
		dll.Append(20)
		dll.Append(30)
		dll.Append(40)
		dll.Append(50)
		dll.Append(60)

		dll.Rotate(0)

		assert.Equal(t, 10, dll.Head.Data)
	})

	t.Run("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=7 positions", func(t *testing.T) {
		dll := New[int]()
		dll.Append(10)
		dll.Append(20)
		dll.Append(30)
		dll.Append(40)
		dll.Append(50)
		dll.Append(60)

		dll.Rotate(7)

		assert.Equal(t, 10, dll.Head.Data)
	})

	t.Run("should rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=4 positions", func(t *testing.T) {
		dll := New[int]()
		dll.Append(10)
		dll.Append(20)
		dll.Append(30)
		dll.Append(40)
		dll.Append(50)
		dll.Append(60)

		dll.Rotate(4)

		assert.Equal(t, 50, dll.Head.Data)
	})
}
