package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	type person struct {
		Id int
		Name string
		Country string
	}

	zhangxiaoan := person{Id: 1001, Name: "zhangxiaoan", Country: "Chian"}

	tmpl, err := template.ParseFiles("./templates/index.tpl")
	if err != nil {
		fmt.Println("Erro happened.")
	}
	tmpl.Execute(response, zhangxiaoan)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", r)
}
