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

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/index.html"))
		fmt.Println(w, r)
		templ.Execute(w, nil)
	})

	http.HandleFunc("/AI1", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/AI1.html"))
		fmt.Println(w, r)
		templ.Execute(w, nil)
	})

	http.HandleFunc("/AI2", func(w http.ResponseWriter, r *http.Request) {
		templ := template.Must(template.ParseFiles("frontend/AI2.html"))
		fmt.Println(w, r)
		templ.Execute(w, nil)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
