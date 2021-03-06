package list

import (
	"net/http"
	"html/template"

	log "github.com/sirupsen/logrus"
	_ "k8s.io/api/core/v1"
	"pilot/daemon"
	"pilot/models/deploy/board"
	"strconv"
	"github.com/gorilla/mux"
	"pilot/session"
	"github.com/gorilla/context"
)

func ListBoards(response http.ResponseWriter, request *http.Request) {

	d, err := daemon.GetInstance(); if err != nil {
		log.Errorf("Daemon GetInstance err:%v", err)
		return
	}
/*
	type showBoard struct {
		Name string
		BoardName string
		Type string
		Chassis string
		Slot string
		Cpu string
		Image string
	}
	showBoards := []showBoard{}
	err = d.BoardStore.Walk(func(b *board.Board) error {
		brd := showBoard{
			Name: b.ProjName,
			BoardName: b.BoardName,
			Type: b.BoardType,
			Chassis: strconv.FormatInt(b.ChassisNumber, 10),
			Slot: strconv.FormatInt(b.SlotNumber, 10),
			Cpu: strconv.FormatInt(b.CpuNumber, 10),
			Image: b.Image,
		}
		log.Debugf("append board:%v b:%v", brd, b)
		showBoards = append(showBoards, brd)
		return nil
	})
	if err != nil {
		log.Errorf("list Container failed:%v\r\n", err)
		return
	}
*/
	type showBoard struct {
		BoardName string
		Status string
		RunNode string
	}
	userName := context.Get(request, session.CLOUDWARE_USER_KEY).(string)
	allContainers, err := d.CloudwareDriver.ListContainers(userName); if err != nil {
		log.Errorf("ListContainers failed :%v", err)
		return
	}

	showBoards := []showBoard{}
	if allContainers != nil {
		for _, container := range allContainers.Items {
			brd := showBoard{
				BoardName: container.BoardName,
				Status: container.Status,
				RunNode: container.RunNode,
			}
			showBoards = append(showBoards, brd)
		}
	}

	tmpl, err := template.ParseFiles("./templates/list_boards.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}

	type listInfo struct{
		UserName string
		ShowBoards []showBoard
	}
	linfo := listInfo{
		UserName: context.Get(request, session.CLOUDWARE_USER_KEY).(string),
		ShowBoards: showBoards,
	}
	err = tmpl.Execute(response, linfo); if err != nil {
		log.Errorf("execute error :%v", err)
	}
}
func ListDevices(response http.ResponseWriter, request *http.Request) {

	d, err := daemon.GetInstance(); if err != nil {
		log.Errorf("Daemon GetInstance err:%v", err)
		return
	}
	/*
		type showBoard struct {
			Name string
			BoardName string
			Type string
			Chassis string
			Slot string
			Cpu string
			Image string
		}
		showBoards := []showBoard{}
		err = d.BoardStore.Walk(func(b *board.Board) error {
			brd := showBoard{
				Name: b.ProjName,
				BoardName: b.BoardName,
				Type: b.BoardType,
				Chassis: strconv.FormatInt(b.ChassisNumber, 10),
				Slot: strconv.FormatInt(b.SlotNumber, 10),
				Cpu: strconv.FormatInt(b.CpuNumber, 10),
				Image: b.Image,
			}
			log.Debugf("append board:%v b:%v", brd, b)
			showBoards = append(showBoards, brd)
			return nil
		})
		if err != nil {
			log.Errorf("list Container failed:%v\r\n", err)
			return
		}
	*/
	type showDevice struct{
		DeviceName string
		Type string
		CSC string
	}
	userName := context.Get(request, session.CLOUDWARE_USER_KEY).(string)
	allDevices, err := d.CloudwareDriver.ListDevices(userName); if err != nil {
		log.Errorf("ListContainers failed :%v", err)
		return
	}

	showDevices:= []showDevice{}
	if allDevices != nil {
		for _, device := range allDevices.Items {
			dev := showDevice{
				DeviceName: device.DeviceName,
				Type: device.Type,
				CSC: device.CSC,
			}
			showDevices = append(showDevices, dev)
		}
	}

	tmpl, err := template.ParseFiles("./templates/list_devices.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}

	type listInfo struct{
		UserName string
		ShowDevices []showDevice
	}
	linfo := listInfo{
		UserName: context.Get(request, session.CLOUDWARE_USER_KEY).(string),
		ShowDevices: showDevices,
	}
	err = tmpl.Execute(response, linfo); if err != nil {
		log.Errorf("execute error :%v", err)
	}
}


func BoardDetails(response http.ResponseWriter, request *http.Request) {
	boardName := mux.Vars(request)["name"]

	d, err := daemon.GetInstance();
	if err != nil {
		log.Errorf("Daemon GetInstance err:%v", err)
		return
	}

	type otherBoard struct {
		BoardName string
	}
	type ifList struct {
		BoardName string
		IfName string
		IfType string
		PeerBoardName string
		PeerIfName string
		OtherBoards []otherBoard
	}

	type showBoard struct {
		Name      string
		BoardName string
		Type      string
		Chassis   string
		Slot      string
		Cpu       string
		Image     string
		IfList    []ifList
		UserName  string
	}

	brd, err := d.BoardStore.Get(boardName)
	if err != nil {
		log.Errorf("get board:%s failed err:%v\r\n", boardName, err)
		return
	}

	otherBoards := []otherBoard{}
	err = d.BoardStore.Walk(func(bd *board.Board) error {
		if bd.BoardName != brd.BoardName {
			otherBoards = append(otherBoards, otherBoard{
				BoardName: bd.BoardName,
			})
		}
		return nil
	})

	if err != nil {
		log.Errorf("Walk error:%v", err)
		return
	}
	iflists := []ifList{}
	for _, ifinter := range brd.BoardInterfaces {
		peerName := ""
		peerBoard := ""
		if ifinter.Endpoint != nil {
			peerName = ifinter.Endpoint.IfName
			peerBoard = ifinter.Endpoint.BoardName
		}
		ifl := ifList{
			BoardName: brd.BoardName,
			IfName: ifinter.IfName,
			IfType: ifinter.IfType,
			PeerBoardName: peerName,
			PeerIfName: peerBoard,
			OtherBoards: otherBoards,
		}
		iflists = append(iflists, ifl)
	}

	sb := showBoard{
		Name: brd.ProjName,
		BoardName: brd.BoardName,
		Type: brd.BoardType,
		Chassis: strconv.FormatInt(brd.ChassisNumber, 10),
		Slot: strconv.FormatInt(brd.SlotNumber, 10),
		Cpu: strconv.FormatInt(brd.CpuNumber, 10),
		Image: brd.Image,
		IfList: iflists,
		UserName: context.Get(request, session.CLOUDWARE_USER_KEY).(string),
	}

	log.Debugf("board detail:%v", sb)
	tmpl, err := template.ParseFiles("./templates/board_details.html", "./templates/bdetail_header.tpl",
	"./templates/navbar.tpl", "./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}
	err = tmpl.Execute(response, sb)
	if err != nil {
		log.Errorf("execute error:%v ", err)
	}
}

func otherIfs() (string, error) {
	return "aaaaa", nil
}
