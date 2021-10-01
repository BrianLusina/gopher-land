package interfaces

import "fmt"

// NilInterface is considered nil when both the type & the value are nil
func NilInterface() {
	var i interface{}
	fmt.Println(i == nil)
	fmt.Printf("%T, %v\n", i, i)

	var s *string
	fmt.Println("s == nil:", s == nil)

	i = s
	// at this point i not nil, because it has a concrete type string. nil interfaces are only considered nil when both
	// the type & the value are nil
	// fmt.Println("i == nil:", i == nil)
	fmt.Printf("%T, %v\n", i, i)
}
