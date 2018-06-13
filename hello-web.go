package main

import (
	"net/http"
	"fmt"
)

func main() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(res, "You requested %s", req.URL.Path)
	})

	//fs := http.FileServer(http.Dir("assets/")) // Where to look for assets and statics
	//http.Handle("/assets/", http.StripPrefix("/assets/", fs))

	//Listen and server i.e server
	http.ListenAndServe(":80", nil)

}
