package session

import (
	"net/http"
	"html/template"
	"fmt"

	"github.com/gorilla/sessions"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/context"
)

const (
	SERCRET="cloudware0.1"
	SESSION_KEY="cloudware"
	CLOUDWARE_USER_KEY="cloudware_user"
)

var store = sessions.NewCookieStore([]byte(SERCRET))

func HomePage(response http.ResponseWriter, request *http.Request) {
	type loginfo struct {
		UserName string
	}
	linfo := loginfo{UserName: context.Get(request, CLOUDWARE_USER_KEY).(string)}
	tmpl, err := template.ParseFiles("./templates/index.tpl","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		logrus.Errorf("Error happened:%v",err)
		return
	}
	tmpl.Execute(response, linfo)
}


func LoginFailed(w http.ResponseWriter) error {
	tmpl, err := template.ParseFiles("./templates/login.html", "./templates/header.tpl",
		"./templates/footer.tpl", "./templates/login_header.tpl")
	if err != nil {
		logrus.Errorf("login failed! :%v", err)
	}
	tmpl.Execute(w, nil)
	return err
}

func GetUserName(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := store.Get(r, SESSION_KEY)
	if err != nil {
		logrus.Errorf("Get session error:%v", err)
		LoginFailed(w)
		return "", fmt.Errorf("invalid user")
	}

	logrus.Debugf("get session:%v", session)
	username := session.Values["username"]
	if username == nil {
		logrus.Errorf("username empty session")
		LoginFailed(w)
		return "", fmt.Errorf("invalid user")

	}
	s := username.(string)
	if s == "" {
		LoginFailed(w)
		return "", fmt.Errorf("invalid user")
	}
	return s, nil
}

func SetUserName(name string, w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, SESSION_KEY)
	if err != nil {
		logrus.Errorf("Set session error:%v", err)
		return err
	}

	session.Values["username"] = name
	session.Save(r, w)
	return nil
}
