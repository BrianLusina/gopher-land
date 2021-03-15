package main

import (
	"fmt"
	f "fmt"
	"runtime"
	"time"

	"github.com/BrianLusina/goffer-land/tour/interfaces"
)

type Distance int

var (
	village Distance = 100
	city             = 50
)

func ofInterfacesAndEmpty() {
	fmt.Println("...")
	fmt.Println("Interfaces...")

	b := interfaces.Book{
		Author:    "David McCullough",
		Name:      "John Adams",
		Published: time.Date(2000, time.May, 22, 0, 0, 0, 0, time.UTC),
	}

	m := interfaces.Movie{
		Name:     "The Goadfather",
		Director: "Frances Ford Coppola",
		Year:     1972,
	}

	fmt.Println(interfaces.Display(b))
	fmt.Println(interfaces.Display(m))

	fmt.Println(interfaces.Empty())

	interfaces.NilInterface()
	fmt.Println("...")
}

func main() {
	fmt.Println("Hello Gopher!")
	f.Println("Hello Gopher!")
	f.Println(runtime.NumCPU())
	f.Println(runtime.Version())

	f.Println(city + int(village))

	ofInterfacesAndEmpty()
}
