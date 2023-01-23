package concept

type shape interface {
	getType() string
	accept(visitor)
}
