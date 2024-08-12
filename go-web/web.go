package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	tf, er := template.ParseFiles("templates/index.html")
	if er != nil {
		tf, _ = template.New("index").Parse("<html><body>body><h1>NO TEMPLATE.</h1></body></html>")
	}

	hh := func(w http.ResponseWriter, r *http.Request) {
		er := tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}

	http.HandleFunc("/", hh)
	http.ListenAndServe(":8080", nil)
}
