package list

import (
	"fmt"
	"net/http"
	"html/template"
	_ "pilot/daemon"
	_ "k8s.io/api/core/v1"
)

func ListBoards(response http.ResponseWriter, request *http.Request) {
	type board struct {
		Name string
		Image string
		Status string
	}
	boards := []board{{Name: "sim01-mpu-0-1-0", Image: "v9trunk:d001", Status:"Up"},
					{Name: "sim01-lpu-0-3-0", Image: "v9trunk:d002", Status:"Down"}}

					/*
	d, err := daemon.GetInstance(); if err != nil {
		fmt.Printf("Daemon GetInstance err:%v", err)
		return
	}

	boards, err := d.ListContainers()
	if err != nil {
		fmt.Printf("list Container failed:%v\r\n", err)
		return
	}

	for _, board := range boards.Items {
		pod := board.(v1.Pod)
	}*/
	tmpl, err := template.ParseFiles("./templates/list_boards.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		fmt.Println("Error happened:%v", err)
		return
	}

	tmpl.Execute(response, boards)
}
