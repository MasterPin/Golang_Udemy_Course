package main

import (
	"html/template"
	"os"
)

func main() {

	t, err := template.New("hello_world").Parse("tmpls/hello_world.tmpl")
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, t)
	if err != nil {
		panic(err)
	}
}
