package singlylinkedlist

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateLinkedList(t *testing.T) {

	t.Run("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 when k=0", func(t *testing.T) {
		sll := New[int]()
		sll.Append(10)
		sll.Append(20)
		sll.Append(30)
		sll.Append(40)
		sll.Append(50)
		sll.Append(60)

		sll.Rotate(0)

		assert.Equal(t, sll.Head.Data, 10)
	})

	t.Run("should not rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=7 positions", func(t *testing.T) {
		sll := New[int]()
		sll.Append(10)
		sll.Append(20)
		sll.Append(30)
		sll.Append(40)
		sll.Append(50)
		sll.Append(60)

		sll.Rotate(7)

		assert.Equal(t, sll.Head.Data, 10)
	})

	t.Run("should rotate a non empty linked list 10 -> 20 -> 30 -> 40 -> 50 -> 60 by k=4 positions", func(t *testing.T) {
		sll := New[int]()
		sll.Append(10)
		sll.Append(20)
		sll.Append(30)
		sll.Append(40)
		sll.Append(50)
		sll.Append(60)

		sll.Rotate(4)

		assert.Equal(t, sll.Head.Data, 50)
	})
}
