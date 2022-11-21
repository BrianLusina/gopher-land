package user

type Collection[T comparable] interface {
	createIterator() Iterator[T]
}
