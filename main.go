package main

import (
	"fmt"
	"net/http"
)

type Strawberry int

func (s Strawberry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}

func main() {
	var s Strawberry
	http.ListenAndServe(":8080", s)
}
