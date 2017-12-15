package login

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"

	"pilot/session"
)

func Registry(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	if method == "POST" {
		request.ParseForm()
		for k, v := range request.Form {
			if k == "username" {
				username := v[0]
				logrus.Debugf("Get username:%s", username)
				err := session.SetUserName(username, response, request)
				if err != nil {
					fmt.Fprintf(response, "Internal server error:%v", err)
					return
				}
			}
		}
	}
}
