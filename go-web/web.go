package main

import (
	"net/http"
)

func main() {
	msg := `<h1>Hello, World!</h1>`
	hh := func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(msg))
	}

	http.HandleFunc("/hello", hh)

	http.ListenAndServe(":8080", nil)
}
