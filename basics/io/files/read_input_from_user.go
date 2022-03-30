package main

import "fmt"

var (
	firstName, lastName, s string
	i                      int
	f                      float32
	input                  = "56.12 / 5212 / Go"
	format                 = "%f / %d / %s"
)

func ReadInputFromUser() {
	fmt.Println("Please enter your full name: ")
	_, err := fmt.Scanln(&firstName, &lastName)
	if err != nil {
		return // or panic()
	}
	fmt.Printf("Hi %s %s!\n", firstName, lastName)
	_, err = fmt.Sscanf(input, format, &f, &i, &s)
	if err != nil {
		return
	}
	fmt.Println("From the string we read: ", f, i, s)
}

func main() {
	ReadInputFromUser()
}
