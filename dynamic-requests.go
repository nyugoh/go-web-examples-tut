package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"fmt"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/students/{name}/{stream}", StudentHandler)

	log.Fatal(http.ListenAndServe(":8080", r))
}

func StudentHandler(w http.ResponseWriter, r *http.Request)  {
	// Get the dynamic parts
	vars := mux.Vars(r)
	name, stream := vars["name"], vars["stream"]

	fmt.Fprintln(w, "Student details")
	fmt.Fprintf(w, "\tName :: %s \n\tStream :: %s", name, stream)
}
