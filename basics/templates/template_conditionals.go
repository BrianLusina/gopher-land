package main

import (
	"fmt"
	"os"
	"text/template"
)

func run_template_cond() {
	tEmpty := template.New("template test 1")
	tEmpty = template.Must(tEmpty.Parse("Empty pipeline if demo: {{if ``}} Will not print. {{end}}\n"))
	err := tEmpty.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	tWithValue := template.New("template test 2")
	tWithValue = template.Must(tWithValue.Parse("Non empty pipeline if demo: {{if `anything`}} Will print. {{end}}\n"))
	err = tWithValue.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}

	tIfElse := template.New("template test 2")
	//non empty pipeline following if condition
	tIfElse = template.Must(tIfElse.Parse("if-else demo: {{if `anything`}} Print IF part. {{else}} Print ELSE part.{{end}}\n"))
	err = tIfElse.Execute(os.Stdout, nil)
	if err != nil {
		fmt.Println("There was an error:", err.Error())
	}
}
