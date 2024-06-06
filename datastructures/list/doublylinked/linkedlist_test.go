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

func TestDeleteNodeByKey(t *testing.T) {
	dll := New[int]()

	dll.Append(1)
	dll.Append(2)
	dll.Append(3)
	dll.Append(4)
	dll.Append(5)

	assert.Equal(t, 5, dll.Size())

	dll.DeleteNodeByKey(3)

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
	type testCase struct {
		data     []any
		expected []any
	}

	testCases := []testCase{
		{
			data:     []any{10, 20, 30},
			expected: []any{30, 20, 10},
		},
		{
			data:     []any{1, 2, 3, 4, 5},
			expected: []any{5, 4, 3, 2, 1},
		},
		{
			data:     []any{"A", "B", "C"},
			expected: []any{"C", "B", "A"},
		},
		{
			data:     []any{2, 4, 6, 7, 5, 1},
			expected: []any{1, 5, 7, 6, 4, 2},
		},
	}

	for _, tc := range testCases {
		testName := fmt.Sprintf("should reverse a linked list of %v to %v", tc.data, tc.expected)
		t.Run(testName, func(t *testing.T) {
			doublyLinkedList := New[any]()
			for _, data := range tc.data {
				doublyLinkedList.Append(data)
			}

			doublyLinkedList.Reverse()

			assert.Equal(t, tc.expected[0], doublyLinkedList.Head.Data)

			current := doublyLinkedList.Head

			actualData := []any{}
			// OR once the rangefunc is no longer behind a feature flag GOEXPERIMENT=rangefunc
			// for node := range doublyLinkedList.All {
			// 	actualData = append(actualData, node.Data)
			// }

			for current != nil {
				actualData = append(actualData, current.Data)
				current = current.Next
			}

			assert.Equal(t, tc.expected, actualData)
		})
	}
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
