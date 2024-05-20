package circularlinked

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testCaseAppendPrepend[T any] struct {
	initialData []T
	data        T
	expected    []T
}

var testCasesAppend = []testCaseAppendPrepend[any]{
	{
		initialData: []any{1, 2, 3, 4, 5, 6},
		data:        7,
		expected:    []any{1, 2, 3, 4, 5, 6, 7},
	},
}

// TestAppend tests adding a new data item to a circular linked list
func TestAppend(t *testing.T) {
	for _, tc := range testCasesAppend {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should append(%v) to get expected={%v}", tc.initialData, tc.data, tc.expected)
		t.Run(testName, func(t *testing.T) {
			linkedList := New[any]()
			for _, data := range tc.initialData {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.initialData[0], actualHead.Data)

			// append new data
			linkedList.Append(tc.data)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.initialData[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}

func BenchmarkAppend(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests in short mode")
	}

	for _, tc := range testCasesAppend {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should append(%v) to get expected={%v}", tc.initialData, tc.data, tc.expected)
		b.Run(testName, func(t *testing.B) {
			linkedList := New[any]()
			for _, data := range tc.initialData {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.initialData[0], actualHead.Data)

			// append new data
			linkedList.Append(tc.data)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.initialData[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}

var testCasesPrepend = []testCaseAppendPrepend[any]{
	{
		initialData: []any{1, 2, 3, 4, 5, 6},
		data:        7,
		expected:    []any{7, 1, 2, 3, 4, 5, 6},
	},
}

// TestPrepend tests adding a new data item to the beginning of a circular linked list
func TestPrepend(t *testing.T) {
	for _, tc := range testCasesPrepend {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should prepend(%v) to get expected={%v}", tc.initialData, tc.data, tc.expected)
		t.Run(testName, func(t *testing.T) {
			linkedList := New[any]()
			for _, data := range tc.initialData {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.initialData[0], actualHead.Data)

			// prepend new data
			linkedList.Prepend(tc.data)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.expected[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}

func BenchmarkPrepend(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests in short mode")
	}

	for _, tc := range testCasesPrepend {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should prepend(%v) to get expected={%v}", tc.initialData, tc.data, tc.expected)
		b.Run(testName, func(t *testing.B) {
			linkedList := New[any]()
			for _, data := range tc.initialData {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.initialData[0], actualHead.Data)

			// append new data
			linkedList.Prepend(tc.data)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.expected[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}

type testCaseDeleteNodeByKey[T any] struct {
	initialData []T
	key         T
	expected    []T
}

var testCasesDeleteNodeByKey = []testCaseDeleteNodeByKey[any]{
	{
		initialData: []any{1, 2, 3, 4, 5, 6},
		key:         4,
		expected:    []any{1, 2, 3, 5, 6},
	},
}

// TestDeleteNodeByKey tests deleting a node by key
func TestDeleteNodeByKey(t *testing.T) {
	for _, tc := range testCasesDeleteNodeByKey {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should delete key(%v) to get expected={%v}", tc.initialData, tc.key, tc.expected)
		t.Run(testName, func(t *testing.T) {
			linkedList := New[any]()
			for _, data := range tc.initialData {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.initialData[0], actualHead.Data)

			// prepend new data
			linkedList.DeleteNodeByKey(tc.key)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.expected[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}

func BenchmarkDeleteNodeByKey(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests in short mode")
	}

	for _, tc := range testCasesDeleteNodeByKey {
		testName := fmt.Sprintf("Initial CircularLinkedList of data={%v} should deleteNodeByKey(%v) to get expected={%v}", tc.initialData, tc.key, tc.expected)
		b.Run(testName, func(t *testing.B) {
			linkedList := New[any]()
			for _, data := range tc.initialData {
				linkedList.Append(data)
			}
			actualHead := linkedList.HeadNode()
			assert.NotNil(t, actualHead)
			assert.Equal(t, tc.initialData[0], actualHead.Data)

			// append new data
			linkedList.DeleteNodeByKey(tc.key)

			actualHead2 := linkedList.HeadNode()
			assert.Equal(t, tc.expected[0], actualHead2.Data)

			actualData := []any{}
			current := actualHead2

			for current != nil {
				actualData = append(actualData, current.Data)

				if current.Next == actualHead2 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expected, actualData)
		})
	}
}
