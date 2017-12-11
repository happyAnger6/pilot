package deploy

import (
	"fmt"
	"net/http"
	"html/template"
	"strconv"

	log "github.com/sirupsen/logrus"

	"pilot/daemon"
	"pilot/deploy/driver"
	"pilot/models/deploy/board"
)

func parseBoard(params map[string][]string)(*board.Board, error) {
	b := &board.Board{}
	for k, v := range params {
		switch k {
		case "bname":
			b.ProjName = v[0]
		case "btype":
			b.BoardType = v[0]
		case "bimage":
			b.Image = v[0]
		case "brunnode":
			b.RunNode = v[0]
		case "bchassis":
			b.ChassisNumber, _ = strconv.ParseInt(v[0], 10, 64)
		case "bslot":
			b.SlotNumber, _ = strconv.ParseInt(v[0], 10, 64)
		case "bcpu":
			b.CpuNumber, _ = strconv.ParseInt(v[0], 10, 64)
		}
	}

	return b, nil
}

func StartBoard(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	log.Debugf("Method:%v", method)
	if method == "POST" {
		request.ParseForm()
		opts := &driver.ContainerOpts{}
		opts.CreateOpts = make(map[string]interface{})
		for k, v  := range request.Form {
			log.Debugf("k:%v\r\n", k)
			log.Debugf("v:%v\r\n", v)
			opts.CreateOpts[k] = v
		}

		d, err := daemon.GetInstance(); if err != nil {
			log.Errorf("Daemon GetInstance err:%v", err)
			return
		}

		bd, err := parseBoard(request.Form); if err != nil {
			log.Errorf("parseBoard err:%v\r\n", err)
			return
		}

		log.Debugf("board:%v\r\n", bd)
		d.BoardStore.Store(bd.ProjName, bd)
		err = d.StartContainer(bd.ProjName, opts)
		if err != nil {
			log.Errorf("start Container failed:%v\r\n", err)
		}
	}
	tmpl, err := template.ParseFiles("./templates/start_board.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}

	tmpl.Execute(response, nil)
}

