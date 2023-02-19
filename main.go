package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Rsvp struct {
	Name  string
	Email string
	Phone string

	WillAttened bool
}

type formData struct {
	*Rsvp
	Errors []string
}

var responses = make([]*Rsvp, 0, 10)
var templates = make(map[string]*template.Template, 3)

func loadTemplates() {
	templatesNames := [5]string{"html-templates/welcome", "html-templates/form", "html-templates/thanks", "html-templates/sorry", "html-templates/list"}

	for index, name := range templatesNames {
		t, err := template.ParseFiles("html-templates/layout.html", name+".html")
		if err == nil {
			templates[name] = t
			fmt.Println("Loaded template", index, name)
		} else {
			panic(err)
		}
	}
}

func welcomeHandler(writer http.ResponseWriter, request *http.Request) {
	templates["html-templates/welcome"].Execute(writer, nil)
}

func listHandler(writer http.ResponseWriter, request *http.Request) {
	templates["html-templates/list"].Execute(writer, nil)
}

func formHandler(writer http.ResponseWriter, request *http.Request) {
	if request.Method == http.MethodGet {
		templates["html-templates/form"].Execute(writer, formData{
			Rsvp: &Rsvp{}, Errors: []string{},
		})
	}
}

func main() {
	loadTemplates()

	http.HandleFunc("/", welcomeHandler)
	http.HandleFunc("/list", listHandler)
	http.HandleFunc("/form", formHandler)

	err := http.ListenAndServe(":5000", nil)
	if err != nil {
		fmt.Println(err)
	}
}
