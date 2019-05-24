package handlers

import (
	"fmt"
	"github.com/gggggggunit/blog/models"
	"net/http"
)

func SavePostH(rw http.ResponseWriter, r *http.Request) {

	//id:=
	//title:= r.FormValue("title")
	//content:= r.FormValue("content")

	err := tpl.ExecuteTemplate(rw, "index", nil)
	if err != nil {
		fmt.Printf("ExecuteTemplate INDEX: %s\n", err)
	}

}
