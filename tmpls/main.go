package main

import (
	"os"
	"text/template"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseFiles("tmpls/hello_world.tmpl"))
}
func main() {

	err := tmpl.ExecuteTemplate(os.Stdout, "hello_world.tmpl", `Release self-focus; embrace other-focus....hocus pocus...liviticus`)
	if err != nil {
		panic(err)
	}
	items := []string{"test1", "test2", "test3", "test4"}
	multVariables := template.Must(template.ParseFiles("tmpls/multiple_variables.tmpl"))
	err = multVariables.ExecuteTemplate(os.Stdout, "multiple_variables.tmpl", items)

	testMap := map[string]int{"42": 42, "one_hundred": 100, "twenty": 20}
	mapParsing := template.Must(template.ParseFiles("tmpls/map.tmpl"))
	err = mapParsing.ExecuteTemplate(os.Stdout, "map.tmpl", testMap)

	name := Names{
		FirstName: "Bob",
		LastName:  "Acri",
	}

	structParse := template.Must(template.ParseFiles("tmpls/struct.tmpl"))
	err = structParse.ExecuteTemplate(os.Stdout, "struct.tmpl", name)

	names := []Names{
		{
			FirstName: "Bob",
			LastName:  "Acri",
		}, {
			FirstName: "Helena",
			LastName:  "Fischer",
		}, {
			FirstName: "Me",
			LastName:  "And You",
		}, {
			FirstName: "LastName",
			LastName:  "FirstName",
		},
	}
	structsParse := template.Must(template.ParseFiles("tmpls/structs.tmpl"))
	err = structsParse.ExecuteTemplate(os.Stdout, "structs.tmpl", names)

}

type Names struct {
	FirstName string
	LastName  string
}
