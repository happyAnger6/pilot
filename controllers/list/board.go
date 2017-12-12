package list

import (
	"net/http"
	"html/template"

	log "github.com/sirupsen/logrus"
	_ "k8s.io/api/core/v1"
	"pilot/daemon"
	"pilot/models/deploy/board"
)

func ListBoards(response http.ResponseWriter, request *http.Request) {
	d, err := daemon.GetInstance(); if err != nil {
		log.Errorf("Daemon GetInstance err:%v", err)
		return
	}

	type showBoard struct {
		Name string
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
			Type: b.BoardType,
			Chassis: string(b.ChassisNumber),
			Slot: string(b.SlotNumber),
			Cpu: string(b.CpuNumber),
			Image: b.Image,
		}
		log.Debugf("append board:%v", brd)
		showBoards = append(showBoards, brd)
		return nil
	})
	if err != nil {
		log.Errorf("list Container failed:%v\r\n", err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/list_boards.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		log.Errorf("Error happened:%v", err)
		return
	}

	tmpl.Execute(response, showBoards)
}
