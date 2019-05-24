package handlers

import (
	"fmt"
	"net/http"
)

func IndexH(rw http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(rw, "index", nil)
	if err != nil {
		fmt.Printf("ExecuteTemplate INDEX: %s\n", err)
	}

}
