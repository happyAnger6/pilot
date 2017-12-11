package deploy

import (
	"net/http"
	"html/template"

	log "github.com/sirupsen/logrus"
)

func CreateTemplate(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	log.Debugf("Method:%v", method)
	if method == "POST" {
		request.ParseForm()
		for k, v  := range request.Form {
			log.Debugf("k:%v\r\n", k)
			log.Debugf("v:%v\r\n", v)
		}
	}
	tmpl, err := template.ParseFiles("./templates/create_template.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v",err)
	}

	tmpl.Execute(response, nil)
}

