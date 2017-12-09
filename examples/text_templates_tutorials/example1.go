package main

import (
	"text/template"
	"os"
)

func main() {
	type Inventory struct {
		Material string
		Count 	 int
	}

	sweaters := Inventory{"wool", 17}
	tmpl, err := template.New("test").Parse("{{.Count}} items are made of {{.Material}}")
	if err != nil {
		panic(err)
	}
	err = tmpl.Execute(os.Stdout, sweaters)
	if err != nil {
		panic(err)
	}
}
