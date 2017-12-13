package deploy

import (
	"net/http"
	"github.com/gorilla/mux"
)

func NetworkConnect(response http.ResponseWriter, request *http.Request) {
	bname := mux.Vars(request)["name"]

}