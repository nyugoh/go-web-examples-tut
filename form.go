package main

import (
	"net/http"
	"log"
	"fmt"
	"html/template"
)

type Message struct {
	Name string
	Subject string
	Body string
}

func main() {

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprintln(writer, "Hello send the message")
	})
	http.HandleFunc("/contact", ContactFormHandler)

	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func ContactFormHandler(writer http.ResponseWriter, request *http.Request) {
	data := struct {
		Success bool
	}{false}
	form, err := template.ParseFiles("views/form.html")
	tmp := template.Must(form, err)
	//todo := template.Must(template.ParseFiles("views/todo.html"))

	if request.Method == http.MethodPost {
		// Send the contact form
		message := Message{
			request.FormValue("name"),
			request.FormValue("subject"),
			request.FormValue("message"),
		}
		data.Success = true
		fmt.Println(message)
	}

	tmp.Execute(writer, data)
}