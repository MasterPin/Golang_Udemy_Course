package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/dog/", dog)
	mux.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", mux)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>You landed on the index Page</h1>")
}

func dog(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Look at this cute dog</h1>")
}

func me(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>I am glad you are here Aleks ;)</h1>")
}
