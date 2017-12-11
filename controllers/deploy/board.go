package deploy

import (
	"fmt"
	"net/http"
	"html/template"
	"pilot/daemon"
	"pilot/deploy/driver"
	"pilot/models/deploy/board"
)

func parseBoard(params map[string][]string)(*board.Board, error) {
	board := &board.Board{}
	for k, v := range params {
		switch k {
		case "bimage":

		}
	}
}
func StartBoard(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	fmt.Printf("Method:%v", method)
	if method == "POST" {
		request.ParseForm()
		opts := &driver.ContainerOpts{}
		for k, v  := range request.Form {
			fmt.Printf("k:%v\r\n", k)
			fmt.Printf("v:%v\r\n", v)
			opts.CreateOpts[k] = v
		}
		d, err := daemon.GetInstance(); if err != nil {
			fmt.Printf("Daemon GetInstance err:%v", err)
			return
		}

		d.BoardStore.Store()
		err = d.StartContainer(opts.CreateOpts["bname"].(string), opts)
		if err != nil {
			fmt.Printf("start Container failed:%v\r\n", err)
		}
	}
	tmpl, err := template.ParseFiles("./templates/start_board.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		fmt.Println("Error happened:%v", err)
		return
	}

	tmpl.Execute(response, nil)
}

