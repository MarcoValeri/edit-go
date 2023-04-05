package main

import (
	"net/http"
	"text/template"
)

func home() {
	tmpl := template.Must(template.ParseFiles("./views/home.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		title := "Hello GO"
		tmpl.Execute(w, title)
	})
}

func main() {
	home()

	http.ListenAndServe(":80", nil)
}
