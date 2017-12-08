package main

import "fmt"
import "net/http"
import "html/template"

func Hello(response http.ResponseWriter, request *http.Request) {
	type person struct {
		Id      int
		Name    string
		Country string
	}

	zhangxiaoan := person{Id: 1001, Name: "zhangxiaoan", Country: "China"}

	tmpl, err := template.ParseFiles("./userall.tpl","./header.tpl","./center.tpl","./footer.tpl")
	if err != nil {
		fmt.Println("Error happened..")
	}
	tmpl.Execute(response,zhangxiaoan)
}

func main() {
	http.HandleFunc("/", Hello)
	http.ListenAndServe(":8080", nil)
}