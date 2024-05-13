package main

import (
	"html/template"
	"net/http"
	"strings"
)

type Course struct {
	ID   int
	Name string
	CH   int
}

type Courses []Course

func ToUpper(s string) string {
	return strings.ToUpper(s)
}

func main() {
	http.HandleFunc("/", templateHandler)
	http.ListenAndServe(":8082", nil)
}

func templateHandler(w http.ResponseWriter, r *http.Request) {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}
	t := template.New("content.html")
	t.Funcs(template.FuncMap{
		"ToUpper": ToUpper,
	})
	t = template.Must(t.ParseFiles(templates...))
	err := t.Execute(w, Courses{
		{1, "Golang", 20},
		{2, "Python", 30},
		{3, "Java", 40},
	})
	if err != nil {
		panic(err)
	}
}
