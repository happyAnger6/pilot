package session

import (
	"github.com/gorilla/sessions"
	"net/http"
	"github.com/sirupsen/logrus"
	"html/template"
)

var store = sessions.NewCookieStore([]byte("cloudware0.1"))

func GetUserName(w http.ResponseWriter, r *http.Request) (string, error) {
	session, err := store.Get(r, "cloudware")
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
	username := session.Values["username"].(string)
	return username, nil
}

func SetUserName(name string, w http.ResponseWriter, r *http.Request) error {
	session, err := store.Get(r, "cloudware")
	if err != nil {
		logrus.Errorf("Set session error:%v", err)
		return err
	}

	session.Values["username"] = name
	session.Save(r, w)
	return nil
}
