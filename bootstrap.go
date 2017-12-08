package main

import (
	"fmt"
	"net/http"
	"html/template"
	"github.com/gorilla/mux"

	"pilot/controllers/deploy"
)

func Hello(response http.ResponseWriter, request *http.Request) {
	type person struct {
		Id int
		Name string
		Country string
	}

	zhangxiaoan := person{Id: 1001, Name: "zhangxiaoan", Country: "Chian"}

	tmpl, err := template.ParseFiles("./templates/index.tpl","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		fmt.Println("Erro happened:%v",err)
	}
	tmpl.Execute(response, zhangxiaoan)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Hello)
	r.HandleFunc("/deploy/create_template", deploy.CreateTemplate)
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", r)
}
