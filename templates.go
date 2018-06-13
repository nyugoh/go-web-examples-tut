package main

import (
	"html/template"
	"log"
	"github.com/gorilla/mux"
	"net/http"
)

type Todo struct {
	Title string
	Done bool
}

type TodoListData struct {
	Todos []Todo
	PageTitle string
}

func main() {
	r := mux.NewRouter()
	data := TodoListData{
		Todos: []Todo{
			{"Task 1", false},
			{"Task 2", true},
			{"Task 3", false},
		},
		PageTitle: "Awesome TODO list",
	}
	layout, err := template.ParseFiles("views/todo.html")

	if err != nil {
		log.Fatal("Template error")
	}

	tmp := template.Must(layout, err)

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		tmp.Execute(w, data)
	})

	// Static assets
	/*
	~ using gorilla/mux
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
	*/

	// ~ using http default Mux
	http.Handle("/assets", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe(":8080", r))
}
