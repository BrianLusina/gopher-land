package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

type Distance int

var (
	village Distance = 100
	city             = 50
)

func main() {
	fmt.Println("Hello Gopher!")
	f.Println("Hello Gopher!")
	f.Println(runtime.NumCPU())
	f.Println(runtime.Version())

	f.Println(city + int(village))
}
