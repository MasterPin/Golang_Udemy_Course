package main

import (
	"html/template"
	"log"
	"net/http"
)

type Strawberry int

var tpl *template.Template

func (s Strawberry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	tpl.ExecuteTemplate(w, "index.gohtml", r.Form)
}

func init() {
	tpl = template.Must(template.ParseFiles("tmpls/index.gohtml"))
}

func main() {
	var s Strawberry
	http.ListenAndServe(":8080", s)
}
