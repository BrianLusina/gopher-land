package user

type Iterator[T comparable] interface {
	hasNext() bool
	next() *T
}
