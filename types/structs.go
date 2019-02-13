package types

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	// assign a field in a struct
	v.X = 9

	fmt.Println(v.X)

	// using a pointer to assign a value to a struct field
	p := &v
	p.X = 1e9 // or (*p).X = 0

	fmt.Println(v)
}

