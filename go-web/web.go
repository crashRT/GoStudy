package main

import (
	"html/template"
	"log"
	"net/http"
)

// Temps is template structure
type Temps struct {
	notemp *template.Template
	indx   *template.Template
	helo   *template.Template
}

// Template for no-template.
func notemp() *template.Template {
	src := "NO TEMPLATE."
	tmp, _ := template.New("index").Parse(src)
	return tmp
}

// setup template function
func setupTemp() *Temps {
	temps := new(Temps)

	temps.notemp = notemp()

	// set index template
	indx, er := template.ParseFiles("templates/index.html")
	if er != nil {
		indx = temps.notemp
	}
	temps.indx = indx

	// set hello template
	helo, er := template.ParseFiles("templates/hello.html")
	if er != nil {
		helo = temps.notemp
	}
	temps.helo = helo

	return temps
}

// index handler
func index(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// hello handler
func hello(w http.ResponseWriter, r *http.Request, tmp *template.Template) {
	er := tmp.Execute(w, nil)
	if er != nil {
		log.Fatal(er)
	}
}

// main program
func main() {
	temps := setupTemp()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		index(w, r, temps.indx)
	})
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		hello(w, r, temps.helo)
	})

	http.ListenAndServe(":8080", nil)

}
