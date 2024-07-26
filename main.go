package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type Strawberry int

var tpl *template.Template

func (s Strawberry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}

	data := struct {
		Method        string
		Host          string
		ContentLength int64
		URL           *url.URL
		Headers       http.Header
		Submissions   url.Values
	}{
		r.Method,
		r.Host,
		r.ContentLength,
		r.URL,
		r.Header,
		r.Form,
	}

	tpl.ExecuteTemplate(w, "index.gohtml", data)
}

func init() {
	tpl = template.Must(template.ParseFiles("tmpls/index.gohtml"))
}

func main() {
	var s Strawberry
	http.ListenAndServe(":8080", s)
}
