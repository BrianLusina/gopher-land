package variables

import (
	"fmt"
)

// declaring variables with type
var (
	name     string
	location string
	age      int
)

// declaring variables with type
var (
	fullName, address string
	phone             int
)

// declaring variables 1 by 1

var code int
var rooms int
var hotelRoom string

// declaring variables with initializers, 1 per variable

var (
	student   string = "Jack Jones"
	teacher   string = "John Doe"
	classroom int    = 1
)

// declaring variables with inferred types when initializer is available

var (
	school     = "Hogwarts"
	classrooms = 50
)

// initialize variables on the same line

var (
	model, price = "JukeBox 90", 9
)

// inside a function := can be used

func test() {
	herName, herLocation := "Eva", "Kiambu"
	herAge := 24
	fmt.Printf("%s (%d) of %s", herName, herAge, herLocation)
}

// variables can contain any type

var action = test

// the type of variable comes last
var c, python, kotlin bool

// variables with initializers with the type declared
var i, j int = 1, 2

func main() {
	// declaring variables with the type inferred
	var js, cpp, java = true, false, "no!"

	var x int
	fmt.Println(x, c, python, kotlin)

	fmt.Println(i, j, js, cpp, java)

	// short variable declarations,
	// only available within functions
	k := 9
	fmt.Println(k)
}
