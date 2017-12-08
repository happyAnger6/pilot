package main

import (
	"fmt"
	"html/template"
	"os"
)

func main() {
	type person struct {
		Id int
		Name string
		Country string
	}

	zhangxiaoan := person{Id: 1001, Name: "zhangxiaoan", Country:"China"}
	fmt.Println("zhangxiaoan = ", zhangxiaoan)

	tmpl := template.New("tmpl")
	tmpl.Parse("Hello {{.Name}} Welcome to go programming...\n")
	tmpl.Execute(os.Stdout, zhangxiaoan)
}
