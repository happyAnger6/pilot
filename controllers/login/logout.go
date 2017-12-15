package login

import (
	"net/http"
	"pilot/session"
)

func Logout(response http.ResponseWriter, request *http.Request) {
	session.SetUserName("", response, request)
	session.LoginFailed(response)
}