package main

import (
	"fmt"
	"net/http"
	"html/template"

	"github.com/gorilla/mux"

	"pilot/controllers/deploy"
	"pilot/controllers/list"
)

func Index(response http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/index.tpl","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		fmt.Println("Error happened:%v",err)
		return
	}
	tmpl.Execute(response, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)

	r.HandleFunc("/deploy/createTemplate", deploy.CreateTemplate)
	r.HandleFunc("/deploy/startBoard", deploy.StartBoard)
	r.HandleFunc("/list/boards", list.ListBoards)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))
	http.ListenAndServe(":8080", r)
}
