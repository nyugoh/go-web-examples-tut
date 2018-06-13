package main

import "net/http"

func main() {
	// Server all the files in the current dir
	http.ListenAndServe(":8080", http.FileServer(http.Dir(".")))
}
