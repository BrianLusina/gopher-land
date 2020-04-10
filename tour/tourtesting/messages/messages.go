package messages

import "fmt"

// Greet function that returns a greeting to passed in name
func Greet(name string) string {
	return fmt.Sprintf("Hello, %v!\n", name)
}

func depart(name string) string {
	return fmt.Sprintf("Goodbye, %v\n", name)
}

