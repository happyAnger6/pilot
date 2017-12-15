package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"github.com/sirupsen/logrus"
	"html/template"
	"fmt"
)

const (
	SERCRET="cloudware0.1"
	SESSION_KEY="cloudware"
)

var store = sessions.NewCookieStore([]byte(SERCRET))

func GetUserName(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := store.Get(r, SESSION_KEY)
	if err != nil {
		logrus.Errorf("Get session error:%v", err)
		tmpl, err := template.ParseFiles("./templates/login.html", "./templates/header.tpl",
						"./templates/footer.tpl")
		if err != nil {
			logrus.Errorf("Error happened:%v",err)
			return "", err
		}
		tmpl.Execute(w, nil)
		return "", nil
	}

	logrus.Debugf("get session:%v", session)
	username := session.Values["username"]
	if username == nil {
		logrus.Errorf("username empty session")
		tmpl, err := template.ParseFiles("./templates/login.html", "./templates/header.tpl",
			"./templates/footer.tpl")
		if err != nil {
			logrus.Errorf("Error happened:%v",err)
			return "", err
		}
		tmpl.Execute(w, nil)
		return "", fmt.Errorf("invalid user")

	}
	return username.(string), nil
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
