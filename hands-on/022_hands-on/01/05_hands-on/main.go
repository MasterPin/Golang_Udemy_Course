package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.Handle("/", http.HandlerFunc(index))
	mux.Handle("/dog/", http.HandlerFunc(dog))
	mux.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", mux)
}

var tpl *template.Template

func index(w http.ResponseWriter, r *http.Request) {
	tpl, err := tpl.ParseFiles("hands-on/golang-web-dev/022_hands-on/01/03_hands-on/template.gohtml")
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(w, "template.gohtml", "This is some Data parsed in through a template file")
	if err != nil {
		log.Fatalln(err)
	}
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Look at this cute dog</h1>")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>I am glad you are here Aleks ;)</h1>")
}
