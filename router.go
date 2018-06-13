package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
)

func main() {
	// New router gorilla/mux package
	r := mux.NewRouter()

	//Handle requests
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello web from Gorilla/mux router.")
	})

	// A dynamic request handler i.e /books/gopher-journey/page/120
	r.HandleFunc("/books/{book}/page/{page}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r) // Extract all the dynamic segments
		fmt.Fprintf(w, "You requested book: %s page # %s", vars["book"], vars["page"])
	})
/*
	// Grouping HandleFunctions
	booksHandle := r.PathPrefix("/books").Subrouter()
	booksHandle.HandleFunc("/", ListBooks)
	booksHandle.HandleFunc("/{title}/page/{page}", getBook)

	//Specifying host
	r.HandleFunc("/books", Archive).Host("www.booksarchive.com")

	// Setting schemes http or https
	r.HandleFunc("/login", LoginPage).Schemes("https")
	r.HandleFunc("/about", AboutPage).Schemes("http")

	// Static assets
	fs := http.FileServer(http.Dir("assets"))
	r.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fs))
*/
	// Start server and pass mux router
	http.ListenAndServe(":80", r)
}
