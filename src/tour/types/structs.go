package types

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	// latitude of the event
	Lat float64

	// Longitude of the event
	Lon float64

	// Date of event
	Date time.Time
}

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	bootamp := Bootcamp{
		Lat: 34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	}

	fmt.Println(bootamp)

	// assign a field in a struct
	v.X = 9

	fmt.Println(v.X)

	// using a pointer to assign a value to a struct field
	p := &v
	p.X = 1e9 // or (*p).X = 0

	fmt.Println(v)
}

