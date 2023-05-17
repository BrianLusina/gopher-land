package types

import "golang.org/x/exp/constraints"

// Number represents all types of numbers, integers or floats
type Number interface {
	constraints.Integer | constraints.Float
}

// String represents all string types
type String interface {
	~string
}

// Comparable types that support ==, <, >
type Comparable interface {
	constraints.Ordered
}
