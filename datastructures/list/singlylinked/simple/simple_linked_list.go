package simple

import "errors"

var (
	ErrEmptyList = errors.New("list is empty")
)

type (
	Element struct {
		Value int
		Next  *Element
	}

	List struct {
		root  *Element
		count int
	}
)

func New(values []int) *List {
	list := &List{}
	for _, value := range values {
		list.Push(value)
	}
	return list
}

func (l *List) Size() int {
	return l.count
}

func (l *List) Push(element int) {
	l.count += 1
	node := &Element{
		Value: element,
		Next:  nil,
	}

	if l.root == nil {
		l.root = node
		return
	}
	current := l.root

	for ; current.Next != nil; current = current.Next {
	}

	current.Next = node
}

func (l *List) Pop() (int, error) {
	if l.root == nil {
		return 0, ErrEmptyList
	}

	current := l.root

	if current.Next == nil {
		data := current.Value
		current = nil
		l.count -= 1
		return data, nil
	}

	for ; current.Next.Next != nil; current = current.Next {
	}

	data := current.Next.Value
	current.Next = nil
	l.count -= 1
	return data, nil
}

func (l *List) Array() []int {
	values := []int{}
	current := l.root
	for ; current != nil; current = current.Next {
		values = append(values, current.Value)
	}
	return values
}

func (l *List) Reverse() *List {
	if l.root == nil {
		return l
	}

	var prev, next *Element
	current := l.root

	for current != nil {
		next = current.Next
		current.Next = prev
		prev = current
		current = next
	}
	l.root = prev
	return l
}
