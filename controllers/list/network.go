package list

import (
	"fmt"
	"net/http"
	"github.com/gorilla/context"
	"pilot/session"
	"pilot/daemon"
	"github.com/sirupsen/logrus"
	"html/template"
)

func ListConnections(response http.ResponseWriter, request *http.Request) {

	userName := context.Get(request, session.CLOUDWARE_USER_KEY).(string)

	d, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("NetworkConnectDevice getinstance failed:%v ", err)
		return
	}

	type showConnection struct{
		LocalDevName string
		LocalPort string
		PeerDevName string
		PeerPort string
	}

	connections, err := d.CloudwareDriver.ListConnections(userName, ""); if err != nil {
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
		ShowConnections []showConnection
	}
	linfo := listInfo{
		UserName: userName,
		ShowConnections: showConnections,
	}

	tmpl, err := template.ParseFiles("./templates/list_connections.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		logrus.Errorf("Error happened:%v", err)
		return
	}
	err = tmpl.Execute(response, linfo); if err != nil {
		logrus.Errorf("execute error :%v", err)
	}
}