package deploy

import (
	"net/http"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"pilot/daemon"
	_ "pilot/models/deploy/board"
	"github.com/gorilla/websocket"
	"github.com/gorilla/context"
	"pilot/session"
	"fmt"
	"html/template"
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

func NetworkConnectDevice(response http.ResponseWriter, request *http.Request) {
	devName := mux.Vars(request)["devName"]
	devType := mux.Vars(request)["devType"]
	devCSC := mux.Vars(request)["devCSC"]
	logrus.Debugf("ConnectDevice name:%s type:%s csc:%s", devName, devType, devCSC)

	method := request.Method
	logrus.Debugf("network connect method:%v", method)

	userName := context.Get(request, session.CLOUDWARE_USER_KEY).(string)

	d, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkConnectDevice getinstance failed:%v ", err)
		return
	}

	if method == http.MethodPost {
		devPort := request.FormValue("bdevport")
		peerDevName := request.FormValue("bpeerdevname")
		peerDevPort := request.FormValue("bpeerdevport")
		logrus.Debugf("Add connection %s-%s to %s-%s ", devName, devPort, peerDevName, peerDevPort)

		err = d.CloudwareDriver.AddConnection(userName, devName, devPort, peerDevName, peerDevPort); if err != nil {
			logrus.Errorf("Add connection err:%v", err)
			fmt.Fprintf(response, "Add connection err:%v", err)
			return
		}
	}

	type showConnection struct{
		LocalDevName string
		LocalPort string
		PeerDevName string
		PeerPort string
	}

	connections, err := d.CloudwareDriver.ListConnections(userName, devName); if err != nil {
		logrus.Errorf("get connections failed:%v", err)
		fmt.Fprintf(response, "get connections failed:%v", err)
		return
	}

	showConnections := []showConnection{}
	if connections != nil {
		for _, connection := range connections.Items {
			showconnection := showConnection {
				LocalDevName: connection.DeviceName,
				LocalPort: connection.PortName,
				PeerDevName: connection.PeerDevice,
				PeerPort: connection.PeerPort,
			}
			showConnections = append(showConnections, showconnection)
		}
	}

	type listInfo struct {
		UserName string
		DeviceName string
		DeviceType string
		DeviceCSC string
		ShowConnections []showConnection
	}
	linfo := listInfo{
		UserName: userName,
		DeviceName: devName,
		DeviceType: devType,
		DeviceCSC: devCSC,
		ShowConnections: showConnections,
	}

	tmpl, err := template.ParseFiles("./templates/list_oneboard_connections.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		logrus.Errorf("Error happened:%v", err)
		return
	}
	err = tmpl.Execute(response, linfo); if err != nil {
		logrus.Errorf("execute error :%v", err)
	}
}

func NetworkDisconnectDevice(response http.ResponseWriter, request *http.Request) {
	devName := mux.Vars(request)["devName"]
	devPort := mux.Vars(request)["devPort"]

	method := request.Method
	logrus.Debugf("network disconnect method:%v", method)

	userName := context.Get(request, session.CLOUDWARE_USER_KEY).(string)

	d, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkDisconnectDevice getinstance failed:%v ", err)
		return
	}

	err = d.CloudwareDriver.RemoveConnection(userName, devName, devPort); if err != nil {
		logrus.Errorf("dis connections failed:%v", err)
		fmt.Fprintf(response, "dis connections failed:%v", err)
		return
	}

	fmt.Fprintf(response, "remove connection successfully!")
	return
}
