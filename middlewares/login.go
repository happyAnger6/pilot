package middlewares

import (
	"net/http"
	"pilot/session"
	"github.com/gorilla/context"
)

func CheckLogin(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		username, err := session.GetUserName(w, r); if err != nil {
			return
		}
		context.Set(r, session.CLOUDWARE_USER_KEY, username)
		f(w, r)
	}
}
