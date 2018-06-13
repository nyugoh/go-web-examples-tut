package main

import (
	"encoding/json"
	"net/http"
	"log"
)

type Book struct {
	Title string `json:"title"`
	Author string `json:"author"`
}

type Books struct {
	Books []Book `json:"books"`
	Total int `json:"total"`
}

func main() {
	goWeb := Book{"Building Web app using Go", "Joe Nyugoh"}
	goWeb1 := Book{"Building Web app using Node.js", "Joe Nyugoh"}
	goWeb2 := Book{"Building Web app using Python", "Joe Nyugoh"}
	 books := Books{}
	 books.Books = append(books.Books, goWeb, goWeb1, goWeb2)
	 books.Total = 3
	book, err := json.MarshalIndent(books, "\n", "  ")
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		if err != nil {
			http.Error(writer, err.Error(), http.StatusInternalServerError)
			return
		}
		writer.Header().Set("Content-Type", "text/json")
		writer.Write(book)
	})

	log.Fatal(http.ListenAndServe(":8000", nil))
}
