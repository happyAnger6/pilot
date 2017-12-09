package main

import (
	"text/template"
	"os"
)

func flag() bool{
	return true
}

func main() {
	type Inventory struct {
		Material string
		Count 	 int
	}

	funcMap := map[string]interface{}{"flag":flag}
	tmpl, err := template.New("test").Funcs(funcMap).Parse("{{if flag}} hello the world!{{end}}")
	if err != nil {
		panic(err)
	}

	err = tmpl.Execute(os.Stdout, nil)
	if err != nil {
		panic(err)
	}
}
