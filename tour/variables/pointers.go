package variables

import "fmt"

func pointers() {
	// initializes a pointer
	var firstName *string = new(string)

	*firstName = "Brian"

	// will print out a memory address
	fmt.Println(firstName)

	// when using a deference operation, the actual value is printed
	fmt.Println(*firstName)
}

func addressOfOperator() {
	lastName := "Lusina"

	// prints out Lusina
	fmt.Println(lastName)

	// using address-of operator
	pointer := &lastName

	// prints out memory address and dereferenced value
	fmt.Println(pointer, *pointer)
}
