package main

import (
	"html/template"
	"fmt"
	"net/http"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	type person struct {
		Id      int
		Name    string
		Country string
	}
	fmt.Println("Hello")
	zhangxiaoan := []person{{Id: 1001, Name: "zhangxiaoan", Country:"China"},
		{Id: 1002, Name: "zhangxiaoan1", Country:"China"}}


	tmpl, err := template.ParseFiles("./tmpl1.html")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response, zhangxiaoan)
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}
