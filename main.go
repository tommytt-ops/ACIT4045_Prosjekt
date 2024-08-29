package main

import (
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

	testHandler := func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/index.html"))
		templ.Execute(w, data)
	}

	addTodohandler := func(w http.ResponseWriter, r *http.Request) {
		message := r.PostFormValue("message")
		templ := template.Must(template.ParseFiles("frontend/index.html"))
		todo := Todo{Id: len(data["Todos"]) + 1, Message: message}
		data["Todos"] = append(data["Todos"], todo)
		templ.ExecuteTemplate(w, "todo-list-element", todo)
	}

	http.HandleFunc("/", testHandler)
	http.HandleFunc("/add-todo", addTodohandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
