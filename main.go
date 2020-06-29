package main

import (
	"fmt"
	f "fmt"
	"runtime"
)

func main() {
	fmt.Println("Hello Gopher!")
	f.Println("Hello Gopher!")
	f.Println(runtime.NumCPU())
	f.Println(runtime.Version())
}
