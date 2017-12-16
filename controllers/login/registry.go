package login

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"pilot/session"
	"pilot/daemon"
)

func Registry(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	var username string
	if method == "POST" {
		request.ParseForm()
		for k, v := range request.Form {
			if k == "username" {
				username = v[0]
				logrus.Debugf("Get username:%s", username)
				err := session.SetUserName(username, response, request)
				if err != nil {
					fmt.Fprintf(response, "Internal server error:%v", err)
					return
				}
			}
		}
	}
	d, err := daemon.GetInstance(); if err != nil {
		logrus.Errorf("Daemon getinstance failed :%v", err)
		fmt.Fprintf(response, "%v", err)
		return
	}

	err = d.UserManagerDriver.AddUser(username); if err != nil {
		logrus.Errorf("AddUser:%s failed :%v", username, err)
		fmt.Fprintf(response, "%v", err)
		return
	}
	session.HomePage(response, request)
}
