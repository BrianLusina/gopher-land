package main

import (
	"fmt"
	"os"
	"text/template"
)

type Person struct {
	Name string
	age  string
}

func run_basic() {
	t := template.New("Person")
	t, _ = t.Parse("hello {{.Name}} you are {{.age}} years old")
	p := Person{Name: "Bob", age: "50"}
	if err := t.Execute(os.Stdout, p); err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
