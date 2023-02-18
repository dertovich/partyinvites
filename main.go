package main

import (
	"fmt"
	"html/template"
)

type Rsvp struct {
	Name  string
	Email string
	Phone string

	WillAttened bool
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templatesNames := [5]string{"welcome", "form", "thanks", "sorry", "list"}

	for index, name := range templatesNames {
		t, err := template.ParseFiles("layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func main() {
	loadTemplates()
}
