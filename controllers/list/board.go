package list

import (
	"fmt"
	"net/http"
	"html/template"
	_ "pilot/daemon"
	_ "k8s.io/api/core/v1"
	"pilot/daemon"
)

func ListBoards(response http.ResponseWriter, request *http.Request) {
	type board struct {
		Name string
		Image string
		Status string
	}
	d, err := daemon.GetInstance(); if err != nil {
		fmt.Printf("Daemon GetInstance err:%v", err)
		return
	}

	boards, err := d.ListContainers()
	if err != nil {
		fmt.Printf("list Container failed:%v\r\n", err)
		return
	}

	tmpl, err := template.ParseFiles("./templates/list_boards.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		fmt.Println("Error happened:%v", err)
		return
	}

	tmpl.Execute(response, boards.Items)
}
