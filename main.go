package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Id      int
	Message string
}

func main() {

	data := map[string][]Todo{
		"Todos": {
			Todo{Id: 1, Message: "Milk"},
		},
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/index.html"))
		fmt.Println(w, r)
		templ.Execute(w, nil)
	})

	http.HandleFunc("/AI1", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/about.html"))
		fmt.Println(w, r)
		templ.Execute(w, nil)
	})

	http.HandleFunc("/AI2", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/form.html"))
		fmt.Println(w, r)
		templ.Execute(w, data)
	})

	http.HandleFunc("/add-todo", func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)
		templ := template.Must(template.ParseFiles("frontend/form.html"))
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
