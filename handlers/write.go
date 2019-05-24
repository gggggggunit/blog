package handlers

import (
	"fmt"
	"net/http"
)

func WriteH(rw http.ResponseWriter, r *http.Request) {

	err := tpl.ExecuteTemplate(rw, "write", nil)
	if err != nil {
		fmt.Printf("ExecuteTemplate WRITE: %s\n", err)
	}
}
