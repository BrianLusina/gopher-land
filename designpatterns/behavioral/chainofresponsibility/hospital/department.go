package hospital

// Department is a Handler interface
type Department interface {
	execute(*Patient)
	setNext(Department)
}
