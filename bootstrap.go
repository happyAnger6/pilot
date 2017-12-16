package main

import (
	"net/http"
	"html/template"

	"github.com/gorilla/mux"
	 log "github.com/sirupsen/logrus"

	"pilot/controllers/deploy"
	"pilot/controllers/list"
	"pilot/session"
	"pilot/controllers/login"
	"pilot/middlewares"
	"github.com/gorilla/context"
)

func Index(response http.ResponseWriter, request *http.Request) {
	type loginfo struct {
		UserName string
	}
	username := context.Get(request, session.CLOUDWARE_USER_KEY).(string)
	linfo := loginfo{UserName: username}
	tmpl, err := template.ParseFiles("./templates/index.tpl","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v",err)
		return
	}
	tmpl.Execute(response, linfo)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", middlewares.CheckLogin(Index))

	r.HandleFunc("/login/registry", login.Registry)
	r.HandleFunc("/logout", login.Logout)

	r.HandleFunc("/deploy/createTemplate", middlewares.CheckLogin(deploy.CreateTemplate))
	r.HandleFunc("/deploy/startBoard", middlewares.CheckLogin(deploy.StartBoard))

	r.HandleFunc("/board/delete/{name}", middlewares.CheckLogin(deploy.DeleteBoard))

	r.HandleFunc("/network/connect", middlewares.CheckLogin(deploy.NetworkConnect))
	r.HandleFunc("/network/disconnect/{name}", middlewares.CheckLogin(deploy.NetworkDisconnect))
	r.HandleFunc("/network/device/connect/{devName}/{devType}/{devCSC}", middlewares.CheckLogin(deploy.NetworkConnectDevice))
	r.HandleFunc("/network/device/disconnect/{devName}/{devPort}", middlewares.CheckLogin(deploy.NetworkDisconnectDevice))

	r.HandleFunc("/list/boards", middlewares.CheckLogin(list.ListBoards))
	r.HandleFunc("/list/devices", middlewares.CheckLogin(list.ListDevices))
	r.HandleFunc("/list/board/details/{name}", middlewares.CheckLogin(list.BoardDetails))

	r.PathPrefix("/").Handler(http.FileServer(http.Dir("static")))

	log.SetLevel(log.DebugLevel)
	log.Debugf("start Listen...\r\n")
	http.ListenAndServe(":8889", r)
}
