package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	t := template.New("test")

	t = template.Must(t.Parse("{{with $3 := `hello`}}{{$3}}{{end}}!\n"))
	err := t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	t = template.Must(t.Parse("{{with $x3 := `hola`}}{{$x3}}{{end}}!\n"))
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	t = template.Must(t.Parse("{{with $x_1 := `hey`}}{{$x_1}} {{.}} {{$x_1}}{{end}}!\n"))
	err = t.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
