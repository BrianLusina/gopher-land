package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t, _ = t.Parse("{{with `hello`}}{{.}}{{end}}!\n")
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}
	t, _ = t.Parse("{{with `hello`}}{{.}} {{with `Mary`}}{{.}}{{end}}{{end}}!\n")
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
