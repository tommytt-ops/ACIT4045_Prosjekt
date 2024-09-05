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

	routehandler := func(w http.ResponseWriter, r *http.Request) {
		// Map URL paths to corresponding template files
		var templateFile string

		// Serve different templates based on the URL path
		switch r.URL.Path {
		case "/":
			templateFile = "frontend/index.html" // Default to index.html for root
		case "/about":
			templateFile = "frontend/about.html" // Serve about.html for the /about URL
		default:
			http.NotFound(w, r) // 404 for any undefined routes
			return
		}

		// Parse and execute the chosen template
		templ := template.Must(template.ParseFiles(templateFile))
		if err := templ.Execute(w, nil); err != nil {
			http.Error(w, "Error rendering template", http.StatusInternalServerError)
		}
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
	//http.HandleFunc("/", routehandler)       // Handle root and other pages like /about
	http.HandleFunc("/about", routehandler)

	log.Fatal(http.ListenAndServe(":8000", nil))
}
