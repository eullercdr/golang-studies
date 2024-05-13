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

type Courses []Course

func main() {
	t := template.Must(template.New("template.html").ParseFiles("template.html"))
	err := t.Execute(os.Stdout, Courses{
		{1, "Golang", 20},
		{2, "Python", 30},
		{3, "Java", 40},
	})
	if err != nil {
		panic(err)
	}
}
