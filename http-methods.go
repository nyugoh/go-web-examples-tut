package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"log"
	"fmt"
)

func main() {
	r := mux.NewRouter()

	// Handle specific http request
	r.HandleFunc("/blog/{blogId}", AddBlog).Methods("POST")
	r.HandleFunc("/blog/{blogId}", GetBlog).Methods("GET").Host("0.0.0.0")
	r.HandleFunc("/blog/{blogId}", DeleteBlog).Methods("DELETE")
	r.HandleFunc("/blog/{blogId}", UpdateBlog).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8080", r))
}

func AddBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Post method")
}

func GetBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Get method")
}

func DeleteBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Delete method")
}

func UpdateBlog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Update method")
}
