package types

import "golang.org/x/exp/constraints"

type Number interface {
	constraints.Integer | constraints.Float
}

// Comparable types that support ==, <, >
type Comparable interface {
	constraints.Ordered
}
