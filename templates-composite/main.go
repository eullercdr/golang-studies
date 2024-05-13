package main

import (
	"net/http"
	"text/template"
)

type Course struct {
	ID   int
	Name string
	CH   int
}

type Courses []Course

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
	t := template.Must(template.New("content.html").ParseFiles(templates...))
	err := t.Execute(w, Courses{
		{1, "Golang", 20},
		{2, "Python", 30},
		{3, "Java", 40},
	})
	if err != nil {
		panic(err)
	}
}
