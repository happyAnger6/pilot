package deploy

import (
	"net/http"
	"html/template"
	"strconv"

	log "github.com/sirupsen/logrus"

	"pilot/daemon"
	"pilot/deploy/driver"
	"pilot/models/deploy/board"
	"github.com/gorilla/mux"
	"pilot/session"
	"github.com/sirupsen/logrus"
	"github.com/gorilla/context"
	"pilot/controllers/list"
)

type loginfo struct {
	UserName string
}

func addBoardInterface(board *board.Board, ifter *board.BoardInterface) error {
	board.BoardInterfaces = append(board.BoardInterfaces, ifter)
	return nil
}

func parseBoard(params map[string][]string)(*board.Board, error) {
	b := &board.Board{}
	var chassis, slot, cpu string
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
			chassis = v[0]
			b.ChassisNumber, _ = strconv.ParseInt(chassis, 10, 64)
		case "bslot":
			slot = v[0]
			b.SlotNumber, _ = strconv.ParseInt(slot, 10, 64)
		case "bcpu":
			cpu = v[0]
			b.CpuNumber, _ = strconv.ParseInt(cpu, 10, 64)
		case "ginter":
			b.GInterfaceNum, _ = strconv.ParseInt(v[0], 10, 64)
		case "tenginter":
			b.TGInterfaceNum, _ = strconv.ParseInt(v[0], 10, 64)
		}
	}

	b.BoardName = b.ProjName + "-" + b.BoardType + "-" + chassis + slot + cpu
	for i := 0; i < int(b.GInterfaceNum); i++ {
		ifiter := &board.BoardInterface{
			BoardName: b.BoardName,
			IfType: board.GigabitEthernet,
			IfName: board.GigabitEthernet + chassis + "/" + slot + "/" + strconv.Itoa(i+1),
		}
		addBoardInterface(b, ifiter)
	}
	for i := 0; i < int(b.TGInterfaceNum); i++ {
		ifiter := &board.BoardInterface{
			BoardName: b.BoardName,
			IfType: board.TenGigabitEthernet,
			IfName: board.TenGigabitEthernet + chassis + "/" + slot + "/" + strconv.Itoa(i+1),
		}
		addBoardInterface(b, ifiter)
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

		username := context.Get(request, session.CLOUDWARE_USER_KEY).(string)
		bd.BoardName = username + bd.BoardName

		log.Debugf("board:%v\r\n", bd)
		d.BoardStore.Store(bd.BoardName, bd)
		opts.CreateOpts["username"] = username
		err = d.CloudwareDriver.StartContainer(username, bd.ProjName, bd.BoardType,
					strconv.FormatInt(bd.ChassisNumber,10), strconv.FormatInt(bd.SlotNumber, 10),
					strconv.FormatInt(bd.CpuNumber, 10))
		if err != nil {
			log.Errorf("start Container failed:%v\r\n", err)
		}
	}
	tmpl, err := template.ParseFiles("./templates/start_board.html","./templates/start_board_header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}

	linfo := loginfo{
		UserName: context.Get(request, session.CLOUDWARE_USER_KEY).(string),
	}
	err = tmpl.Execute(response, linfo); if err != nil {
		logrus.Debugf("Execute err:%v", err)
	}
}

func DeleteBoard(response http.ResponseWriter, request *http.Request) {
	name := mux.Vars(request)["name"]
	log.Debugf("DeleteBoard name:%v",  name)

	d, err := daemon.GetInstance(); if err != nil {
		log.Errorf("Daemon GetInstance err:%v", err)
		return
	}

	d.BoardStore.Delete(name)
	username := context.Get(request, session.CLOUDWARE_USER_KEY).(string)
	err = d.CloudwareDriver.RemoveContainer(username, name)
	if err != nil {
		log.Errorf("remove Container failed:%v\r\n", err)
	}

	list.ListBoards(response, request)
	/*
	tmpl, err := template.ParseFiles("./templates/list_boards.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}

	linfo := loginfo{
		UserName: username,
	}
	tmpl.Execute(response, linfo)*/
}
