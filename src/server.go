package main

import (
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	http.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request){
		io.WriteString(w, "Testing!")
		// io.WriteString(w, w)
		})
	http.ListenAndServe(":8000", nil)
}

