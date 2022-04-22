package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t := template.New("test")
	t = template.Must(t.Parse("{{with $x := `hello`}}{{printf `%s %s` $x `Mary`}}{{end}}!\n"))
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
