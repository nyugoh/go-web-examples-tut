package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	r := mux.NewRouter()

	// Handle specific http request
	r.HandleFunc("/blog/{blogId}", AddBlog).Methods("POST")
	r.HandleFunc("/blog/{blogId}", GetBlog).Methods("GET")
	r.HandleFunc("/blog/{blogId}", DeleteBlog).Methods("DELETE")
	r.HandleFunc("/blog/{blogId}", UpdateBlog).Methods("PUT")
}

func AddBlog(w http.ResponseWriter, r *http.Request)  {

}

func GetBlog(w http.ResponseWriter, r *http.Request)  {

}

func DeleteBlog(w http.ResponseWriter, r *http.Request)  {

}

func UpdateBlog(w http.ResponseWriter, r *http.Request)  {

}