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

	tmpl, err := template.ParseFiles("./tmpl1.html")
	if err != nil {
		fmt.Println("Error happend..")
	}
	tmpl.Execute(os.Stdout, zhangxiaoan)
}
