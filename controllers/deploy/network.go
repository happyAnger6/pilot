package deploy

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"pilot/daemon"
	"pilot/models/deploy/board"
)

func NetworkConnect(response http.ResponseWriter, request *http.Request) {
	bname := mux.Vars(request)["bname"]
	ifname := mux.Vars(request)["ifname"]

	logrus.Debugf("Network connect board name:%s ifname:%s", bname, ifname)

	d, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkConnect getinstance failed:%v ", err)
		return
	}

	brd, err := d.BoardStore.Get(bname); if err != nil {
		logrus.Errorf("get board:%s info error:%v", bname, err)
		return
	}

	binters := brd.BoardInterfaces
	binter := &board.BoardInterface{}
	for _, inter := range binters {
		if inter.IfName == ifname {
			binter = inter
		}
	}

	logrus.Debugf("network connect interface:%v", binter)
}

func NetworkDisconnect(response http.ResponseWriter, request *http.Request) {
	bname := mux.Vars(request)["name"]
	logrus.Debugf("Network connect board name:%s", bname)

	_, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkConnect getinstance failed:%v ", err)
		return
	}



}
