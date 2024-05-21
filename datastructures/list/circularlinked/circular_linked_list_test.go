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

type testCaseLength[T any] struct {
	data     []T
	expected int
}

var testCasesLength = []testCaseLength[any]{
	{
		data:     []any{1, 2, 3, 4, 5, 6},
		expected: 6,
	},
}

// TestLength tests getting the length of a linked list
func TestLength(t *testing.T) {
	for _, tc := range testCasesLength {
		testName := fmt.Sprintf("CircularLinkedList of data={%v} should return length of expected={%v}", tc.data, tc.expected)
		t.Run(testName, func(t *testing.T) {
			linkedList := New[any]()
			for _, data := range tc.data {
				linkedList.Append(data)
			}

			actual := linkedList.Length()
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func BenchmarkLength(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests in short mode")
	}

	for _, tc := range testCasesLength {
		testName := fmt.Sprintf("CircularLinkedList of data={%v} should return length of expected={%v}", tc.data, tc.expected)
		b.Run(testName, func(t *testing.B) {
			linkedList := New[any]()
			for _, data := range tc.data {
				linkedList.Append(data)
			}

			actual := linkedList.Length()
			assert.Equal(b, tc.expected, actual)
		})
	}
}

type testCaseSplitList[T any] struct {
	data        []T
	expectedOne []T
	expectedTwo []T
}

var testCasesSplit = []testCaseSplitList[any]{
	{
		data:        []any{1, 2, 3, 4, 5, 6},
		expectedOne: []any{1, 2, 3},
		expectedTwo: []any{4, 5, 6},
	},
}

// TestSplitList tests splitting a list into two halves
func TestSplitList(t *testing.T) {
	for _, tc := range testCasesSplit {
		testName := fmt.Sprintf("CircularLinkedList of data={%v} should split into %v and %v", tc.data, tc.expectedOne, tc.expectedTwo)
		t.Run(testName, func(t *testing.T) {
			circularLinkedList := CircularLinkedList[any]{}
			for _, data := range tc.data {
				circularLinkedList.Append(data)
			}

			actualOne, actualTwo := circularLinkedList.SplitList()
			assert.NotNil(t, actualOne)
			assert.NotNil(t, actualTwo)

			actualHead1 := actualOne.HeadNode()
			assert.Equal(t, tc.expectedOne[0], actualHead1.Data)

			actualData1 := []any{}
			current := actualHead1

			for current != nil {
				actualData1 = append(actualData1, current.Data)

				if current.Next == actualHead1 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(t, tc.expectedOne, actualData1)

			actualHead2 := actualTwo.HeadNode()
			assert.Equal(t, tc.expectedTwo[0], actualHead2.Data)

			actualData2 := []any{}
			current1 := actualHead2

			for current1 != nil {
				actualData2 = append(actualData2, current1.Data)

				if current1.Next == actualHead2 {
					break
				}
				current1 = current1.Next
			}

			assert.ElementsMatch(t, tc.expectedTwo, actualData2)
		})
	}
}

func BenchmarkSplitList(b *testing.B) {
	if testing.Short() {
		b.Skipf("Skipping benchmark tests in short mode")
	}

	for _, tc := range testCasesSplit {
		testName := fmt.Sprintf("CircularLinkedList of data={%v} should split into %v and %v", tc.data, tc.expectedOne, tc.expectedTwo)
		b.Run(testName, func(b *testing.B) {
			circularLinkedList := CircularLinkedList[any]{}
			for _, data := range tc.data {
				circularLinkedList.Append(data)
			}

			actualOne, actualTwo := circularLinkedList.SplitList()
			assert.NotNil(b, actualOne)
			assert.NotNil(b, actualTwo)

			actualHead1 := actualOne.HeadNode()
			assert.Equal(b, tc.expectedOne[0], actualHead1.Data)

			actualData1 := []any{}
			current := actualHead1

			for current != nil {
				actualData1 = append(actualData1, current.Data)

				if current.Next == actualHead1 {
					break
				}
				current = current.Next
			}

			assert.ElementsMatch(b, tc.expectedOne, actualData1)

			actualHead2 := actualTwo.HeadNode()
			assert.Equal(b, tc.expectedTwo[0], actualHead2.Data)

			actualData2 := []any{}
			current1 := actualHead2

			for current1 != nil {
				actualData2 = append(actualData2, current1.Data)

				if current1.Next == actualHead2 {
					break
				}
				current1 = current1.Next
			}

			assert.ElementsMatch(b, tc.expectedTwo, actualData2)
		})
	}
}
