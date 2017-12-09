package deploy

import (
	"fmt"
	"net/http"
	"html/template"
)

func CreateTemplate(response http.ResponseWriter, request *http.Request) {
	method := request.Method
	fmt.Printf("Method:%v", method)
	if method == "POST" {
		request.ParseForm()
		for k, v  := range request.Form {
			fmt.Printf("k:%v\r\n", k)
			fmt.Printf("v:%v\r\n", v)
		}
	}
	tmpl, err := template.ParseFiles("./templates/create_template.html","./templates/header.tpl",
		"./templates/navbar.tpl","./templates/footer.tpl")
	if err != nil {
		fmt.Println("Error happened:%v",err)
	}

	tmpl.Execute(response, nil)
}

