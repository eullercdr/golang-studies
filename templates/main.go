package main

import (
	"os"
	"text/template"
)

type Course struct {
	ID   int
	Name string
	CH   int
}

func main() {
	course := Course{
		ID:   1,
		Name: "Golang",
		CH:   20,
	}
	tpl := template.New("CourseTemplate")
	tpl, _ = tpl.Parse("ID: {{.ID}} Name: {{.Name}} CH: {{.CH}}")
	err := tpl.Execute(os.Stdout, course)
	if err != nil {
		panic(err)
	}
}
