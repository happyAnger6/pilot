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
	b := &board.Board{}
	for k, v := range params {
		switch k {
		case "bimage":
			b.Image = v[0]
		case "bname":
			b.ProjName = v[0]
		}
	}

	return b, nil
}

func StartBoard(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	fmt.Printf("Method:%v", method)
	if method == "POST" {
		request.ParseForm()
		opts := &driver.ContainerOpts{}
		opts.CreateOpts = make(map[string]interface{})
		for k, v  := range request.Form {
			fmt.Printf("k:%v\r\n", k)
			fmt.Printf("v:%v\r\n", v)
			opts.CreateOpts[k] = v
		}

		d, err := daemon.GetInstance(); if err != nil {
			fmt.Printf("Daemon GetInstance err:%v", err)
			return
		}

		bd, err := parseBoard(request.Form); if err != nil {
			fmt.Printf("parseBoard err:%v\r\n", err)
			return
		}

		fmt.Printf("board:%v\r\n", bd)
		d.BoardStore.Store(bd.ProjName, bd)
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

