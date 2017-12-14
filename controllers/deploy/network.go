package deploy

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"pilot/daemon"
	_ "pilot/models/deploy/board"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func NetworkConnect(response http.ResponseWriter, request *http.Request) {
	conn, err := upgrade.Upgrade(response, request, nil)
	if err != nil {
		logrus.Errorf("get conn err:%v", err)
		return
	}

	type connInfo struct{
		BoardName string
		IfName string
	}
	cinfo := &connInfo{}
	err = conn.ReadJSON(cinfo); if err != nil {
		logrus.Errorf("ReadJson failed:%v", err)
		return
	}

	logrus.Debugf("readjson obj:%v", cinfo)
}
/*func NetworkConnect(response http.ResponseWriter, request *http.Request) {
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
}*/

func NetworkDisconnect(response http.ResponseWriter, request *http.Request) {
	bname := mux.Vars(request)["name"]
	logrus.Debugf("Network connect board name:%s", bname)

	_, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkConnect getinstance failed:%v ", err)
		return
	}



}
