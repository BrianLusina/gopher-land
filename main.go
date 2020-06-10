package main

import "fmt"
import f "fmt"
import "runtime"

func main() {
	fmt.Println("Hello Gopher!")
	f.Println("Hello Gopher!")
	f.Println(runtime.NumCPU())
}
