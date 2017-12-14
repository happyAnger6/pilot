package main

import (
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
	 log "github.com/sirupsen/logrus"

	"pilot/controllers/deploy"
	"pilot/controllers/list"
)

func Index(response http.ResponseWriter, request *http.Request) {

	tmpl, err := template.ParseFiles("./templates/index.tpl","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v",err)
		return
	}
	tmpl.Execute(response, nil)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", Index)

	r.HandleFunc("/deploy/createTemplate", deploy.CreateTemplate)
	r.HandleFunc("/deploy/startBoard", deploy.StartBoard)

	r.HandleFunc("/board/delete/{name}", deploy.DeleteBoard)

	r.HandleFunc("/network/connect", deploy.NetworkConnect)
	r.HandleFunc("/network/disconnect/{name}", deploy.NetworkDisconnect)

	r.HandleFunc("/list/boards", list.ListBoards)
	r.HandleFunc("/list/board/details/{name}", list.BoardDetails)

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.SetLevel(log.DebugLevel)
	log.Debugf("start Listen...\r\n")
	http.ListenAndServe(":8080", r)
}
