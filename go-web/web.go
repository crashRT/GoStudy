package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	html := `
	<!DOCTYPE html>
	<html>
		<head>
			<title>Web</title>
		</head>
		<body>
			<h1>Hello World!</h1>
			<p> Go server is running.</p>
		</body>
	</html>`

	tf, er := template.New("web").Parse(html)
	if er != nil {
		log.Fatal(er)
	}

	hh := func(w http.ResponseWriter, r *http.Request) {
		er = tf.Execute(w, nil)
		if er != nil {
			log.Fatal(er)
		}
	}

	http.HandleFunc("/", hh)
	http.ListenAndServe(":8080", nil)
}
