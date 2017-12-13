package deploy

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"pilot/daemon"
)

func NetworkConnect(response http.ResponseWriter, request *http.Request) {
	bname := mux.Vars(request)["name"]
	logrus.Debugf("Network connect name:%s", bname)

	d, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkConnect getinstance failed:%v ", err)
		return
	}


}